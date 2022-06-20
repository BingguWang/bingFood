package prod

import (
    "time"
)

type Sku struct {
    SkuId      uint64     `json:"skuId" gorm:"column:sku_id;primarykey"` // 单品ID
    SkuName    string     `json:"skuName" gorm:"sku_name"`               // sku名称
    ProdId     uint64     `json:"prodId"`                                // 商品ID
    Price      int        `json:"price" gorm:"price"`                    // 价格
    Weight     int        `json:"weight" gorm:"weight"`                  // 份量，单位克
    SellStatus uint8      `json:"sellStatus" gorm:"sell_status"`         // 是否  售罄0 未售罄 1 售罄
    Stock      int        `json:"stock" gorm:"stock"`                    // sku库存
    CreateAt   time.Time  `json:"createAt" gorm:"autoCreateTime"`        // 创建时间
    UpdateAt   time.Time  `json:"updateAt" gorm:"autoUpdateTime"`        // 最近更新时间
    DeleteAt   *time.Time `json:"deleteAt"`                              // 删除时间，软删除
}

func (*Sku) TableName() string {
    return "t_sku" // 返回你要自定义的表名
}
