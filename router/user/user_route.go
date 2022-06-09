package user

import (
    "github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
    group := r.Group("/user")
    group.POST("/getCode", GetValidCode())
    group.POST("/register", RegisterUserMiddleware())
    group.POST("/login", LoginUserMiddleware())
}