package middleware

import (
    "bingFood/entity/order"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func AddOrderMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        order := &order.Order{}
        ctx.ShouldBind(&order)

        fmt.Println(order)
        // doSomething。。。

        ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
    }
}
