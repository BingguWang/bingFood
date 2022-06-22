package resp

import (
    "bingFood/entity/order"
)

type SettleOrderResp struct {
    ShopId         uint64 // 商家id
    UserId         uint64 // 用户
    UserMobile     string // 用户手机号
    ReceiverMobile string // 接收人号码

    ProdName string // 逗号拼接，产品名称
    ProdNums int    // 商品数量

    OrderStatus   uint8 `json:"orderStatus,omitempty" `  // 订单状态
    DeleteStatus  uint8 `json:"deleteStatus,omitempty"`  // 订单删除状态  0：没有删除， 1：回收站， 2：永久删除
    PayStatus     uint8 `json:"payStatus,omitempty"`     // 支付状态
    RefundStatus  uint8 `json:"refundStatus,omitempty"`  // 订单退款状态
    DeliverStatus uint8 `json:"deliverStatus,omitempty"` // 订单配送状态

    PackingAmount  int // 打包费用
    DeliverAmount  int // 配送费
    ProdAmount     int // 仅商品总价值
    DiscountAmount int // 优惠金额
    FinalAmount    int // 最终支付金额

    DeliverType uint8 // 配送方式，1 外卖配送 2 到店自提

    OrderItems []order.OrderItem // 订单项
}
