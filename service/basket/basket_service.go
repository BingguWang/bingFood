package basket

import (
    "bingFood/entity/basket"
    "bingFood/entity/basket/req"
    "bingFood/global"
    "bingFood/utils"
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
