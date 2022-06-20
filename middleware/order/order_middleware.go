package order

import (
    "bingFood/entity/order"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

// 新增Order
func AddOrderMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        order := &order.Order{}
        ctx.ShouldBind(&order)

        fmt.Println(order)
        // doSomething。。。

        ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
    }
}
