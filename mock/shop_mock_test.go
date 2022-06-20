package mock

import (
    "bingFood/entity/shop/req"
    "bingFood/utils"
    "fmt"
    "testing"
)

func TestShop(t *testing.T)  {
    var shop req.ShopReq

    fmt.Println(utils.ToJsonString(shop))

}
