package prod

import (
    "time"
)

/**
  属性
*/
type Property struct {
    PropId   uint64 `json:"propId" gorm:"primarykey"` // 属性id
    PropName string `json:"propName" gorm:""`         // 属性名称
    ShopId   string `json:"shopId" gorm:""`           // 店铺id
    ProdId   uint64 `json:"prodId"`                   // 商品id

    CreateAt time.Time  `json:"createAt" gorm:"autoCreateTime"` // 创建时间
    UpdateAt time.Time  `json:"updateAt" gorm:"autoUpdateTime"` // 最近更新时间
    DeleteAt *time.Time `json:"deleteAt"`                       // 删除时间，软删除
}

func (*Property) TableName() string {
    return "t_prod_prop" // 返回你要自定义的表名
}
