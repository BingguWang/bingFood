package order

import (
    "bingFood/entity/order"
    "bingFood/entity/order/req"
    "bingFood/entity/order/resp"
    "bingFood/entity/pay"
    "bingFood/global"
    "bingFood/utils"
    "fmt"
    "github.com/pkg/errors"
    "gorm.io/gorm"
    "log"
    "time"
)

/**
微信支付需要的参数
outTradeNo----支付单号，就是settleNo
TotalFee----总共支付了多少钱
userId----用户的唯一标识

支付回到里的有用信息：
outTradeNo----支付单号，就是settleNo
resultCode----返回码，"SUCCESS"


*/
func PayOrder(param req.PayOrderReq) (interface{}, error) {
    db := global.MYSQL_DB
    var od []order.Order
    if err := db.Where(&order.Order{OrderNumber: param.OrderNumber}).
        Where("order_status = 0").
        Find(&od).Error; err != nil {
        return nil, err
    }

    if len(od) == 0 {
        return nil, errors.New(fmt.Sprintf("传入订单号不正确，不存在此未支付的订单:%v", param.OrderNumber))
    }
    oder := od[0]
    var (
        outTradeNo string  // 支付单号
        totalFee   int     // 一共支付多少钱
        openid     = "123" // 用户唯一标识
    )

    number := global.SNOW_NODE.Generate() //
    outTradeNo = number.String()

    orderPay := order.OrderPay{
        OrderNumber: oder.OrderNumber,
        ShopId:      oder.ShopId,
        UserId:      oder.UserId,
        UserMobile:  oder.UserMobile,
        PayNo:       outTradeNo,
        PayAmount:   totalFee,
        PayTypeName: "微信支付",
        PayType:     1,
        PayStatus:   0, // 回到成功后才修改为1
    }
    if err := InsertOrderPay(orderPay); err != nil {
        return nil, err
    }

    // 模拟调用wx支付接口...
    // ...
    result := MockWxPay(totalFee, outTradeNo, openid)

    return resp.PayOrderResp{WxPayMpOrderResult: result}, nil
}

/**
模拟微信支付:传入参数
outTradeNo----支付单号，就是settleNo
TotalFee----总共支付了多少钱
openid----用户的唯一标识

*/


func MockWxPay(totalFee int, outTrade, userId string) pay.WxPayMpOrderResult {
    time.Sleep(500 * time.Millisecond)
    return pay.WxPayMpOrderResult{}
}

func InsertOrderPay(odPay order.OrderPay) error {
    log.Printf("orderPay is : %v", utils.ToJsonString(odPay))

    db := global.MYSQL_DB
    if err := db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(&odPay).Error; err != nil {
            log.Printf("insert orderPay failed : %v", err.Error())
            return err
        }
        return nil
    }); err != nil {
        return err
    }
    return nil
}
