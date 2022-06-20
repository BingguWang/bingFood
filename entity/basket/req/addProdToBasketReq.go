package req

type AddProdToBasketReq struct {
    BasketId uint64 `json:"basket_id;"`
    ShopId   int    `json:"shopId"`   //  店铺id
    UserId   uint64 `json:"userId"`   // 用户id
    SkuId    uint64 `json:"skuId"`    // 单品ID
    ProdNums int    `json:"prodNums"` // 商品数量
}
