package req

import "bingFood/common/request"

type ListProdReq struct {
	PageInfo request.PageInfo `json:"pageInfo"`

	ShopId int `json:"shopId"`
}
