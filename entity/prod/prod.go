package prod

import (
    "time"
)

type Prod struct {
    Id             uint64 `json:"prodId" gorm:"column:prod_id;primary_key;` // 商品id
    ProdName       string `json:"prodName"`                                 // 商品名称
    ShopId         int    `json:"shopId"`                                   //  店铺id
    OriPrice       int    `json:"oriPrice"`                                 // 原价
    PackingFee     int    `json:"packingFee"`                               // 打包费
    Price          int    `json:"price"`                                    // 现价
    Pic            string `json:"pic"`                                      // 商品主图
    Imgs           string `json:"imags"`                                    // 商品图片,分割
    Description    string `json:"description"`                              // 简要描述,卖点等
    Content        string `json:"content"`                                  // 套餐内容
    ProdStatus     uint8  `json:"prodStatus"`                               // 默认是1示正常状态, -1表示删除, 0下架
    IsAutoAdd      uint8  // 默认是0，是否自动补充库存，即库存是否无限制 0不是 1 是
    SoldHistoryNum int    `json:"soldHistoryNum"` // 销量
    SoldMonthNum   int    `json:"soldMonthNum"`   // 月销量
    TotalStock     int    `json:"totalStock"`     // 总库存
    Score          int    `json:"score"`          // 可得积分
    BuyLimit       int    `json:"buyLimit"`       // COMMENT限购数量 0 不限制 大于0表示限制数量

    TodaySoldOut uint8 `gorm:"-"` // 默认是0，表示今日未售罄, 1今日售罄, 是否售罄放到redis里

    CreateAt time.Time `json:"createAt" gorm:"autoUpdateTime"` // 创建时间
    UpdateAt time.Time `json:"updateAt" gorm:"autoUpdateTime"` // 最近更新时间

    CategoryId uint16 `json:"categoryId"` // 商品分类

    Properties []Property `json:"properties" gorm:"foreignKey:ProdId"`
    Skus       []Sku      `json:"skus" gorm:"foreignKey:ProdId"`
}

func (*Prod) TableName() string {
    return "t_prod" // 返回你要自定义的表名
}
