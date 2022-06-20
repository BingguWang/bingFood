package basket

import (
    "bingFood/middleware/basket"
    "github.com/gin-gonic/gin"
)

func BasketRouter(r *gin.Engine) {
    group := r.Group("/basket")
    group.POST("/detail", basket.GetBasketDetail())
    group.POST("/prod/add", basket.AddProdToBasket())
}
