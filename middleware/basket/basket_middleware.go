package basket

import (
    "bingFood/common/response"
    "bingFood/entity/basket"
    "bingFood/entity/basket/req"
    "bingFood/global"
    bs "bingFood/service/basket"
    "bingFood/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "log"
)

func GetBasketDetail() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var basketReq req.BasketReq
        _ = ctx.ShouldBindJSON(&basketReq)

        pageSize := basketReq.PageInfo.PageSize
        page := basketReq.PageInfo.Page

        res, total, err := bs.GetBasketDetails(basketReq)
        if err != nil {
            response.FailWithMessage(fmt.Sprintf("获取购物车列表失败, err:%v", err.Error()), ctx)
            ctx.Abort()
            return
        }
        response.OkWithDetailed(response.PageResult{
            List:     res,
            Total:    total,
            Page:     page,
            PageSize: pageSize,
        }, "获取列表成功", ctx)
        ctx.Abort()
        return
    }
}

func AddProdToBasket() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var reqParam req.AddProdToBasketReq
        _ = ctx.ShouldBindJSON(&reqParam)

        if reqParam.ProdNums <= 0 {
            _ = DeleteBasketRecord(reqParam)
            response.OkWithMessage("商品已移出购物车", ctx)
            ctx.Abort()
            return
        }

        if err := InsertBasketRecord(reqParam); err != nil {
            log.Printf("insert basket failed : %v", err.Error())
            response.FailWithMessage(err.Error(), ctx)
            ctx.Abort()
            return
        }
        response.OkWithMessage("添加商品到购物车成功", ctx)
        ctx.Abort()
        return
    }
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
