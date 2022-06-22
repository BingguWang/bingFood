package order

import (
    "bingFood/common/response"
    "bingFood/entity/order"
    "bingFood/entity/order/req"
    "bingFood/global"
    ord "bingFood/service/order"
    "bingFood/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/pkg/errors"
    "gorm.io/gorm"
    "log"
    "time"
)

func PayOrderMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {

        var payOrderReq req.PayOrderReq
        ctx.ShouldBindJSON(&payOrderReq)

        res, err := ord.PayOrder(payOrderReq)
        if err != nil {
            response.FailWithMessage(err.Error(), ctx)
            ctx.Abort()
            return
        }
        response.OkWithDetailed(res, "支付成功", ctx)
        ctx.Abort()
        return
    }
}

func NoticePayOrderMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var noticeReq req.NoticePayOrderReq
        ctx.ShouldBindJSON(&noticeReq)

        fmt.Println(utils.ToJsonString(noticeReq))
        if noticeReq.XmlData == "SUCCESS" {
            if err := PaySuccess(noticeReq.PayNo); err != nil {
                response.FailWithMessage(err.Error(), ctx)
                ctx.Abort()
                return
            }
        }

        response.OkWithMessage("支付成功", ctx)
        ctx.Abort()
        return
    }
}

func PaySuccess(payNo string) error {
    log.Printf("payNo is : %v", payNo)
    db := global.MYSQL_DB

    if err := db.Transaction(func(tx *gorm.DB) error {
        var orderPay []order.OrderPay
        if err := db.Where(&order.OrderPay{PayNo: payNo}).Find(&orderPay).Error; err != nil {
            return err
        }
        if len(orderPay) == 0 {
            return errors.New(fmt.Sprintf("支付信息有误,payNo:%v", payNo))
        }

        // 修改订单支付表的支付状态
        log.Printf("修改订单支付表信息...")
        tx2 := tx.Model(&order.OrderPay{}).Where("pay_no = ? AND pay_status = 0", payNo).
            Update("pay_status", 1)
        if rows := tx2.RowsAffected; rows == 0 {
            return errors.New(fmt.Sprintf("the orderPay has been paid , payNo : %v", payNo))
        }

        // 修改订单状态为已支付
        log.Printf("修改订单表信息...")
        txx := tx.Model(&order.Order{}).Where("order_number = ? AND order_status = 0", orderPay[0].OrderNumber).
            Select("order_status", "pay_type", "pay_at").
            Updates(map[string]interface{}{"order_status": 1, "pay_type": orderPay[0].PayType, "pay_at": time.Now()})
        if rows := txx.RowsAffected; rows == 0 {
            return errors.New(fmt.Sprintf("order has been paid , orderNumber : %v", orderPay[0].OrderNumber))
        }

        return nil
    }); err != nil {
        return err
    }

    return nil
}
