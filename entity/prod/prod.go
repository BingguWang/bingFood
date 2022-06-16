package prod

import "time"

type Prod struct {
    ProdId         uint64 `gorm:"prod_id"`     // 商品id
    ProdName       string `gorm:"prod_name"`   // 商品名称
    ShopId         int    `gorm:"shop_id"`     //  店铺id
    OriPrice       int    `gorm:"ori_price"`   // 原价
    PackingFee     int    `gorm:"packing_fee"` // 打包费
    Price          int    `gorm:"price"`       // 现价
    Pic            string `gorm:"pic"`         // 商品主图
    Imags          string `gorm:"imgs"`        // 商品图片,分割
    Description    string `gorm:"description"` // 简要描述,卖点等
    Content        string `gorm:"content"`     // 套餐内容
    ProdStatus     uint8  `gorm:"prod_status"` // 默认是1示正常状态, -1表示删除, 0下架
    IsAutoAdd      uint8  // 默认是0，是否自动补充库存，即库存是否无限制 0不是 1 是
    SoldHistoryNum int    `gorm:"sold_history_num"` // 销量
    SoldMonthNum   int    `gorm:"sold_month_num"`   // 月销量
    TotalInventory int    `gorm:"total_inventory"`  // 总库存
    Score          int    `gorm:"score"`            // 可得积分
    BuyLimit       int    `gorm:"buy_limit"`        // COMMENT限购数量 0 不限制 大于0表示限制数量

    TodaySoldOut uint8 // 默认是0，表示今日未售罄, 1今日售罄, 是否售罄放到redis里

    CreateAt time.Time `gorm:"create_at"` // 创建时间
    UpdateAt time.Time `gorm:"update_at"` // 最近更新时间

    CategoryId uint16 `gorm:"category_id"` // 商品分类

}
