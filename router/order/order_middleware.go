package order

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    order2 "orderModule/entity/order"
)

func AddOrderMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        order := &order2.Order{}
        ctx.ShouldBind(&order)

        fmt.Println(order)
        // doSomething。。。

        ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
    }
}
