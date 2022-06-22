package order

import (
    "bingFood/common/response"
    "bingFood/entity/order/req"
    ordsev "bingFood/service/order"
    "github.com/gin-gonic/gin"
    "log"
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

        od, err := ordsev.SettleOrder(ctx, reqParam)
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
        var reqParam req.ConfirmOrderReq
        if err := ctx.ShouldBindJSON(&reqParam); err != nil {
            panic(err)
        }

        orderNumber, err := ordsev.ConfirmOrder(ctx, reqParam)
        if err != nil {
            log.Printf("提交订单失败 : %v", err.Error())
            response.FailWithMessage("提交订单失败"+err.Error(), ctx)
            ctx.Abort()
            return
        }
        response.OkWithDetailed(orderNumber, "提交订单成功", ctx)
        ctx.Abort()
        return
    }
}
