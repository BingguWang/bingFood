package prod
/**
    属性
 */
type Property struct {
    PropId   uint64 `gorm:"prop_id"`   // 属性id
    PropName string `gorm:"prop_name"` // 属性名称
    ShopId   string `gorm:"shop_id"`   // 店铺id
    ProdId   uint64 `gorm:"prod_id"`   // 商品id
}
