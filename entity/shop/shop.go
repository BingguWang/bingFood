package shop

import "time"

type Shop struct {
    ShopId   uint64 `gorm:"column:shop_id"`   // 店铺id
    ShopName string `gorm:"column:shop_name"` // 店铺名称
    Intro    string `gorm:"column:intro"`     // 店铺简介
    ShopLogo string `gorm:"column:shop_logo"` // 店铺logo(可修改)

    MobileBackgroundPic string `gorm:"column:mobile_background_pic"` // 店铺移动端背景图

    ShopStatus uint8 `gorm:"column:shop_status"` // 店铺状态(-1:已删除 0: 停业中 1:营业中,因为我这没有写后台系统所以就是直接默认是营业的)

    Type     uint8     `gorm:"type"`                     // 店铺类型
    CreateAt time.Time `gorm:"create_at;autoCreateTime"` // 创建时间
    UpdateAt time.Time `gorm:"update_at;autoUpdateTime"` // 最近更新时间
}

func (*Shop) TableName() string {
    return "t_shop"
}
