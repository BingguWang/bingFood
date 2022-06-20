package basket

import (
    "bingFood/entity/basket"
    "bingFood/entity/basket/req"
    "bingFood/global"
    "bingFood/utils"
    "fmt"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "log"
)

func GetBasketDetails(param req.BasketReq) (res interface{}, total int64, err error) {
    log.Printf("request args are : %v", utils.ToJsonString(param))

    db := global.MYSQL_DB
    var basketList []basket.Basket
    limit := param.PageInfo.PageSize
    offset := param.PageInfo.Page

    db.Where(&param.BasketCond).Count(&total)
    err = db.Limit(limit).Offset(offset).Where(&param.BasketCond).Find(&basketList).Error
    if err != nil {
        return
    }
    return basketList, total, nil
}



func InsertBasketRecord(param req.AddProdToBasketReq) error {
    log.Printf("request args are : %v", utils.ToJsonString(param))

    db := global.MYSQL_DB
    basketRow := basket.Basket{
        BasketId: param.BasketId,
        ShopId:   param.ShopId,
        SkuId:    param.SkuId,
        UserId:   param.UserId,
        ProdNums: param.ProdNums,
    }

    if err := db.Transaction(func(tx *gorm.DB) error {
        err := tx.Clauses(clause.OnConflict{
            Columns:   []clause.Column{{Name: "shop_id"}, {Name: "user_id"}, {Name: "sku_id"}}, // key column，如果id已存在则变为更新操作
            DoUpdates: clause.AssignmentColumns([]string{"sku_id", "prod_nums"}),               // 更新操作要更新的字段,更新为新值
        }).Create(&basketRow).Error
        return err
    }); err != nil {
        log.Printf("insert prod failed : %v", err.Error())
        return err
    }
    return nil
}
// TODO 有空自己封装一个复制值接口,或直接用copier
func DeleteBasketRecord(param req.AddProdToBasketReq) error {
    log.Printf("request args are : %v", utils.ToJsonString(param))

    db := global.MYSQL_DB
    var deletedRow = basket.Basket{
        BasketId: param.BasketId,
        UserId:   param.UserId,
        ShopId:   param.ShopId,
        SkuId:    param.SkuId,
    }
    fmt.Println(utils.ToJsonString(deletedRow))
    if err := db.Transaction(func(tx *gorm.DB) error {
        e := tx.Where(&deletedRow).Delete(&deletedRow).Error
        return e
    }); err != nil {
        log.Printf("delete row failed : %v", err.Error())
        return err
    }
    return nil
}
