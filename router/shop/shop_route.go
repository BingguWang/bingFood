package shop

import (
    "bingFood/middleware/shop"
    "github.com/gin-gonic/gin"
)

func ShopRouter(r *gin.Engine) {
    group := r.Group("/shop")
    {
        group.POST("/list", shop.GetShopList())
        group.POST("/detail", shop.GetShopDetail())

    }
}
