package user

import "time"

/**
  用户配送地址
*/
type UserDeliveryAddr struct {
    UserDeliveryAddrId int    `json:"userDeliveryAddrId;" gorm:"primaryKey"`
    UserId             uint64 `json:"userId"`
    Receiver           string `json:"receiver"`   // 接收人名称
    ProvinceId         int    `json:"provinceId"` // 省id
    Province           string `json:"province"`   // 省名称
    CityId             int    `json:"cityId"`     // 城市id
    City               string `json:"city"`       // 城市名称
    AreaId             int    `json:"areaId"`     // 区id
    Area               string `json:"area"`       // 区名称
    Detail             string `json:"detail"`     // 详细地址

    CreateAt time.Time `json:"createAt" gorm:"autoCreateTime"` // 创建时间
    UpdateAt time.Time `json:"updateAt" gorm:"autoUpdateTime"` // 修改时间
}
