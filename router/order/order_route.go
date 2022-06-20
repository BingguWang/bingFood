package order

import (
    "bingFood/middleware/order"
    "github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.Engine) {
    group := r.Group("/order")
    group.POST("/add", order.AddOrderMiddleware())
}