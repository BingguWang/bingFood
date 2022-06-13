package req

type UserLoginOrRegisterParam struct {
    UserMobile string `json:"userMobile"` // 手机号
    LoginType  uint8  `json:"loginType"`  // 登录方式 0 手机号登录 1 密码登录 2 第三方微信登录
    ValidCode  string `json:"validCode"`  // 验证码
    Password   string `json:"password"`   // 登录密码，账号密码登录不需要验证码
    //IsLogin    uint8  `json:"isLogin"`    // 1登录 2注册
}
