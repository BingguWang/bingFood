package delivery

import "time"

// 配送单
type DeliveryOrder struct {
    DeliverId       uint64 // 配送单id
    DeliverNumber   string // 配送单号,雪花算法生成
    DeliverymanId   string // 骑手id
    DeliverymanName string // 骑手名称
    DeliveryFee     int    // 配送费
    DeliveryStatus  int    // 0 未配送 1 配送中 2 配送完成

    ReceiveTime time.Time // 送达时间
    Timeout     time.Time // 超时多少时间

    IsGood  uint   // 1 满意 2 还行 3 不满意
    Comment string // 评语
}
