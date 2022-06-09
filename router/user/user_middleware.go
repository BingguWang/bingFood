package user

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "orderModule/common/response"
    "orderModule/entity/user"
    "orderModule/entity/user/req"
    "orderModule/utils"
)

func RegisterUserMiddleware() gin.HandlerFunc { // gin.HandlerFunc其实就是func(*context)
    return func(ctx *gin.Context) {
        user := &req.UserLoginOrRegisterParam{}
        ctx.ShouldBind(&user)

        fmt.Println(user)
        // doSomething。。。

        //validCode := utils.SendMsg(user.UserMobile)
        //response.OkWithData(validCode, ctx)
    }
}
func GetValidCode() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        param := &req.UserLoginOrRegisterParam{}
        ctx.ShouldBind(param)
        validCode, err := utils.SendMsg(param.UserMobile)
        if err != nil {
            response.FailWithMessage(err.Error(), ctx)
        } else {
            response.OkWithData(validCode, ctx)
        }
    }
}
func LoginUserMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        user := &user.User{}
        ctx.ShouldBind(&user)

        fmt.Println(user)
        // doSomething。。。

        ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
    }
}
