package order

import (
    "bingFood/middleware"
    "bingFood/middleware/order"
    "github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.Engine) {
    group := r.Group("/order", middleware.JWTAuthMiddleware())
    group.POST("/toSettle", order.ToSettleOrderMiddleware())
    group.POST("/confirm", order.ConfirmOrderMiddleware())
    group.POST("/pay", order.PayOrderMiddleware())
    group.POST("/pay/notice/wechat", order.NoticePayOrderMiddleware()) // 微信支付回调
}
