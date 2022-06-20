package user

import (
    "bingFood/middleware"
    "bingFood/middleware/user"
    "github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
    group := r.Group("/user")
    {
        group.POST("/getCode", user.GetValidCode())
        group.POST("/getCaptcha", middleware.GetCaptcha())
        group.POST("/verifyCaptcha", middleware.VerifyCaptcha())
        group.POST("/register", user.LoginOrRegUserMiddleware())
        group.POST("/login", user.LoginOrRegUserMiddleware())

        CommonGroup := group.Group("/api")
        CommonGroup.Use(middleware.JWTAuthMiddleware())
        {
            CommonGroup.POST("/userList", user.GetUserList())
        }
    }

}
