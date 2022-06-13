package user

import (
    "bingFood/middleware"
    "github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
    group := r.Group("/user")
    {
        group.POST("/getCode", middleware.GetValidCode())
        group.POST("/getCaptcha", middleware.GetCaptcha())
        group.POST("/verifyCaptcha", middleware.VerifyCaptcha())
        group.POST("/register", middleware.LoginOrRegUserMiddleware())
        group.POST("/login", middleware.LoginOrRegUserMiddleware())

        CommonGroup := group.Group("/api")
        CommonGroup.Use(middleware.JWTAuthMiddleware())
        {
            CommonGroup.POST("/userList", middleware.GetUserList())
        }
    }

}
