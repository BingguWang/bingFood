package middleware

import (
    "bingFood/common/response"
    "bingFood/utils"
    "fmt"
    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenString := ctx.Request.Header.Get("token")

        if tokenString == "" {
            response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", ctx)
            ctx.Abort()
            return
        }
        fmt.Println(tokenString)

        claims, err := utils.ParseToken(tokenString)
        if err != nil {
            if err == utils.TokenIsExpired {
                response.FailWithDetailed(gin.H{"reload": true}, "token授权已过期", ctx)
                ctx.Abort()
                return
            }
            response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
            ctx.Abort()
            return
        }
        ctx.Set("claims", claims)
        ctx.Next()
    }
}