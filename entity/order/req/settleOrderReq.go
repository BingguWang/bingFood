package req

type SettleOrderReq struct {
    BasketIds []uint64 `json:"basket_ids"`
}
