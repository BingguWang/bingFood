package order

type OrderItem struct {
    OrderItemId uint64 // 订单项id
    ShopId      int    // 商家id
    OrderNumber string // 订单号
    ProdId      uint64 // 商品id
    ProdName    string // 商品名称
    ProdNum     int    // 商品个数
    ProdAmount  int    // 商品总价
    Score       int    // 此订单项拥有的积分

    IsCommented uint   // 是否评价 0 未评价 1 已评价
    IsGood      uint   // 1 好评 2 差评
    Comment     string // 评语
}