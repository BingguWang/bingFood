package prod

import "time"

type Sku struct {
    SkuId      uint64    `gorm:"sku_id"`      // 单品ID
    SkuName    string    `gorm:"sku_name"`    // sku名称
    ProdId     uint64    `gorm:"prod_id"`     // 商品ID
    Price      int       `gorm:"price"`       // 价格
    Weight     int       `gorm:"weight"`      // 份量，单位克
    SellStatus uint8     `gorm:"sell_status"` // 是否  售罄0 未售罄 1 售罄
    Stock      int       `gorm:"stock"`       // sku库存
    CreateAt   time.Time `gorm:"create_at"`   // 创建时间
    UpdateAt   time.Time `gorm:"update_at"`   // 最近更新时间
    DeleteAt   time.Time `gorm:"delete_at"`   // 删除时间，软删除
}
