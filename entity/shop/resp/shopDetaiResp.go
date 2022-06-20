package resp

import (
    "bingFood/entity/prod"
    "bingFood/entity/shop"
)

type ShopDetailResp struct {
    Shop shop.Shop
    ProdList []prod.Prod
}