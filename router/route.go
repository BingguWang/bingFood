package router

import (
    "github.com/gin-gonic/gin"
    "orderModule/router/order"
    "orderModule/router/user"
)

type RouteWork func(*gin.Engine)

var routeWorkSlice = []RouteWork{}

func SetupRouter() *gin.Engine {
    r := gin.Default()
    injectRouteWork(user.UserRouter, order.OrderRouter)
    for _, work := range routeWorkSlice {
        work(r)
    }
    return r
}

// 注入路由工作
func injectRouteWork(rwork ...RouteWork) {
    routeWorkSlice = append(routeWorkSlice, rwork...)
}
