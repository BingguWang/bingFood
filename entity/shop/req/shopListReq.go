package req

import (
    "bingFood/common/request"
    "bingFood/entity/shop"
)

type ShopReq struct {
    PageInfo request.PageInfo `json:"pageInfo"`

    ShopCond shop.Shop `json:"shopCond,omitempty"`
}
