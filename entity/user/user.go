package user

import "time"

type User struct {
    UserId            string    `db:"user_id"` // 雪花算法生成
    UserMobile        string    // 用户手机号，这个其实就作为用户名
    LoginPassword     string    // 用户密码
    UserMail          string    // 用户邮箱
    UserWxNum         string    // 微信号
    UserScores        int       // 用户拥有积分
    UserNickName      string    // 用户昵称
    UserRealName      string    // 用户真实姓名
    UserBirthDate     string    // 用户生日
    UserRegRegion     string    // 用户注册所在地区
    UserRegIp         string    // 用户注册所在IP
    LastLoginRegion   string    // 上次登录所在地区
    LastLoginIp       string    // 上次登录所在IP
    UserSex           uint8     // 用户性别
    UserStatus        uint8     // 用户状态
    CreateAt          time.Time // 创建时间
    UpdateAt          time.Time // 修改时间
    UserLastLoginTime time.Time // 用户上次登录时间
    Score             int       // 用户拥有的积分
}
