package delivery

import "time"

// 这些主要应该是在配送模块的，另外搞个服务
type Deliveryman struct {
    DeliverymanId       string
    DeliverymanName     string    // 骑手名
    DeliverymanMail     string    // 骑手邮箱
    DeliverymanMobile   string    // 骑手手机号
    DeliverymanWxNum    string    // 微信号
    DeliverymanRealName string    // 骑手真实姓名
    DeliverymanRegIp    string    // 骑手注册所在IP
    DeliverymanSex      uint8     // 骑手性别
    DeliverymanStatus   uint8     // 骑手状态
    LoginPassword       string    // 骑手登录密码
    CreateAt            time.Time // 创建时间
    UpdateAt            time.Time // 修改时间
}
