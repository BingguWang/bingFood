package req

import (
    "bingFood/common/request"
    "bingFood/entity/basket"
)

type BasketReq struct {
    PageInfo request.PageInfo `json:"pageInfo,omitempty"`

    BasketCond basket.Basket `json:"basketCond,omitempty"`
}
