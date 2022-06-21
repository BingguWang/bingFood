package req

type SettleOrderReq struct {
    BasketIds []uint64 `json:"basketIds"`
}
