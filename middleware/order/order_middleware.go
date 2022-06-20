package order

import (
    "bingFood/common/response"
    "bingFood/entity/basket"
    "bingFood/entity/order"
    "bingFood/entity/order/req"
    "bingFood/global"
    "bingFood/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "strconv"
    "time"
)

func ToSettleOrderMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var reqParam req.SettleOrderReq
        ctx.ShouldBind(&reqParam)
        if reqParam.BasketIds == nil || len(reqParam.BasketIds) == 0 {
            response.FailWithMessage("传入的购物车没有商品", ctx)
            ctx.Abort()
            return
        }

        od, err := SettleOrder(ctx, reqParam)
        if err != nil {
            log.Printf("request args are : %v", err.Error())
            response.FailWithMessage("结算出错", ctx)
            ctx.Abort()
            return
        }
        response.OkWithDetailed(od, "结算order成功", ctx)
        ctx.Abort()
        return
    }
}

// 提交订单
func ConfirmOrderMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {

    }
}

func SettleOrder(ctx *gin.Context, param req.SettleOrderReq) (res interface{}, err error) {
    log.Printf("request args are : %v", utils.ToJsonString(param))

    db := global.MYSQL_DB
    var basketList []basket.Basket

    if err = db.Where("basket_id IN ? ", param.BasketIds).Preload("Sku").Find(&basketList).Error; err != nil {
        return
    }

    fmt.Println(utils.ToJsonString(basketList))

    var (
        oriPriceTotal   int // 原价总和
        packingFeeTotal int // 打包费
        priceTotal      int // 现价总和
        finalTotal      int // 最后需支付金额
        discountTotal   int // 总共优惠的金额
        deliverFeeTotal int // 配送费
        redPacket       int // 红包
        itemList        []order.OrderItem
        prodNums        int // 总商品个数
        shopId          uint64
    )
    for _, v := range basketList {
        sku := v.Sku
        item := order.OrderItem{
            UserId:     v.UserId,
            ShopId:     v.ShopId,
            ProdId:     0,
            ProdName:   sku.ProdName,
            ProdNums:   v.ProdNums,
            Pic:        sku.Pic,
            Price:      sku.Price,
            ProdAmount: sku.Price * v.ProdNums,
            OriPrice:   sku.OriPrice,
            SkuId:      sku.SkuId,
            SkuName:    sku.SkuName,
            PropId:     sku.ProdId,
            PropName:   sku.ProdName + sku.SkuName,
        }
        oriPriceTotal += sku.OriPrice
        priceTotal += sku.Price
        packingFeeTotal += sku.PackingFee

        prodNums += v.ProdNums
        itemList = append(itemList, item)
        shopId = v.ShopId
    }

    // TODO 配送费应该从配送系统计算得到，这里只是用个数值替一下
    deliverFeeTotal = 5 * 100 // 假设是固定的配送费

    discountTotal = (oriPriceTotal - priceTotal) + redPacket
    finalTotal = packingFeeTotal + priceTotal + deliverFeeTotal - discountTotal

    claims, _ := ctx.Get("claims")
    userClaims := claims.(*utils.UserClaims)
    orderRes := order.Order{
        ShopId:         shopId,
        UserMobile:     userClaims.UserMobile,
        ProdNums:       prodNums,
        PackingAmount:  packingFeeTotal,
        DeliverAmount:  deliverFeeTotal,
        ProdAmount:     priceTotal,
        DiscountAmount: discountTotal,
        FinalTotal:     finalTotal,
        OrderItems:     itemList,
    }

    fmt.Println(utils.ToJsonString(orderRes))

    // 返回的结算内容存到redis里,后面的提交订单时不需要前端再传过来了,提交订单的时候删掉
    cli := global.GVA_REDIS
    key := "settledOrder_" + strconv.FormatUint(shopId, 10) + "_" + userClaims.UserMobile // TODO 规范,常数写到其他地方去
    _, err = cli.Set(ctx, key, utils.ToJsonString(orderRes), 10*time.Minute).Result()     // 停留在结算页面没操作超过10分钟结算就作废
    if err != nil {
        return
    }

    return orderRes, nil
}
