package router

import (
    "bingFood/router/order"
    "bingFood/router/prod"
    "bingFood/router/shop"
    "bingFood/router/user"
    "github.com/gin-gonic/gin"
)

type RouteWork func(*gin.Engine)

var routeWorkSlice = []RouteWork{}

func SetupRouter() *gin.Engine {
    r := gin.Default()
    injectRouteWork(user.UserRouter, order.OrderRouter, prod.ProdRouter, shop.ShopRouter)
    for _, work := range routeWorkSlice {
        work(r)
    }
    return r
}

// 注入路由工作
func injectRouteWork(rwork ...RouteWork) {
    routeWorkSlice = append(routeWorkSlice, rwork...)
}
