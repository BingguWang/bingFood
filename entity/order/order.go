package order

import (
    "orderModule/entity/user"
    "time"
)

type Order struct {
    OrderId     uint64
    OrderNumber string // 订单号，雪花算法生成
    ShopId      int    // 商家id
    UserId      string // 用户
    DeliverNumber   string // 配送单号

    ProdName string // 逗号拼接，产品名称
    ProdNum  int    // 商品数量

    OrderStatus   uint8 // 订单状态
    RefundStatus  uint8 // 订单退款状态
    DeliverStatus uint8 // 订单配送状态

    PackingAmount  int // 打包费用
    DeliverAmount  int // 配送费
    ProdAmount     int // 仅商品总价值
    DiscountAmount int // 优惠金额
    FinalTotal     int // 最终支付金额
    Score          int // 本单可得积分

    PayType     uint8 // 支付方式
    DeliverType uint8 // 配送方式，1 外卖配送 2 到店自提

    Remarks string // 备注

    CreateAt         time.Time // 创建时间
    UpdateAt         time.Time // 订单最近更新时间
    PayAt            time.Time // 订单支付时间
    FinishAt         time.Time // 订单完成时间
    CancelAt         time.Time // 订单取消时间
    CancelReasonType int       // 订单取消原因
    DeleteStatus     uint8     // 订单删除状态  0：没有删除， 1：回收站， 2：永久删除

    OrderItems     []OrderItem           // 订单项
    ReceiveAddr    user.UserDeliveryAddr // 接收地址
    ReceiverMobile string                // 接收人号码
}
