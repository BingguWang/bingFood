package order

import (
    "bingFood/middleware"
    "bingFood/middleware/order"
    "github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.Engine) {
    group := r.Group("/order").Use(middleware.JWTAuthMiddleware())
    group.POST("/toSettle", order.ToSettleOrderMiddleware())
    group.POST("/confirm", order.ConfirmOrderMiddleware())
}