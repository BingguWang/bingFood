package req

import (
    "bingFood/entity/user"
)

type ConfirmOrderReq struct {
    Remarks        string                `json:"remarks"`               // 备注
    ReceiveAddr    user.UserDeliveryAddr `json:"receiveAddr,omitempty"` // 接收地址
    ReceiverMobile string                `json:"receiverMobile"`        // 接收人号码
    RedPacket      int                   `json:"redPacket"`             // 红包
    PayStatus      uint8                 `json:"payStatus"`             // 支付状态
    ShopId         uint64                `json:"shopId"`                // 商家id
}
