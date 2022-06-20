package order

import "time"

type OrderItem struct {
    OrderItemId uint64 `gorm:"primaryKey"` // 订单项id
    OrderNumber string // 订单号

    UserId uint64 // 用户
    Score  int    // 此订单项拥有的积分

    ShopId     int    // 商家id
    ProdId     uint64 // 商品id
    ProdName   string // 商品名称
    ProdNums   int    // 商品个数
    Pic        string // 商品图片地址
    ProdAmount int    // 商品总价
    SkuId      uint64 // 产品SkuID
    SkuName    string // 商品sku名称
    Price      int    // 商品单价

    PropId   uint64 // 属性id
    PropName string // 属性名称

    IsCommented uint   // 是否评价 0 未评价 1 已评价
    IsGood      uint   // 1 好评 2 差评 3 一般
    Comment     string // 评语

    CreateAt  time.Time `json:"createAt" gorm:"autoCreateTime"`
    UpdateAt  time.Time `json:"updateAt" gorm:"autoUpdateTime"`
    CommentAt *time.Time
}
