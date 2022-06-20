package order

import (
    "bingFood/entity/user"
    "time"
)

type Order struct {
    OrderId        uint64                `gorm:"primaryKey"`
    OrderNumber    string                // 订单号，雪花算法生成
    ShopId         int                   // 商家id
    UserId         uint64                // 用户
    UserMobile     string                // 用户手机号
    ReceiveAddr    user.UserDeliveryAddr // 接收地址
    ReceiverMobile string                // 接收人号码

    DeliverNumber string // 配送单号
    ProdName      string // 逗号拼接，产品名称
    ProdNums      int    // 商品数量

    OrderStatus   uint8 // 订单状态
    DeleteStatus  uint8 // 订单删除状态  0：没有删除， 1：回收站， 2：永久删除
    PayStatus     uint8 // 支付状态
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

    CreateAt         time.Time `json:"createAt" gorm:"autoCreateTime"` // 创建时间
    UpdateAt         time.Time `json:"updateAt" gorm:"autoUpdateTime"` // 订单最近更新时间
    DeleteAt         *time.Time
    PayAt            *time.Time // 订单支付时间
    FinishAt         *time.Time // 订单完成时间
    CancelAt         *time.Time // 订单取消时间
    CancelApplyAt    *time.Time // 订单申请取消时间
    CancelReasonType uint8      // 订单取消原因

    OrderItems []OrderItem // 订单项
}
