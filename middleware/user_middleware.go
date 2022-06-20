package middleware

import (
    "bingFood/common/response"
    "bingFood/dao"
    "bingFood/entity/user"
    "bingFood/entity/user/req"
    "bingFood/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    uuid "github.com/satori/go.uuid"
)

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

// 未注册的会自动注册
func LoginOrRegUserMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        userParam := &req.UserLoginOrRegisterParam{}
        ctx.ShouldBindJSON(&userParam)
        fmt.Println(userParam)

        if userParam.UserMobile == "" {
            response.FailWithDetailed(gin.H{"reload": true}, "手机号不能为空", ctx)
            ctx.Abort()
            return
        }
        u := user.User{UserMobile: userParam.UserMobile}

        ret := dao.GetUserByCond(u)
        if len(ret) == 0 {
            Register(ctx, userParam)
        } else {
            Login(ctx, userParam)
        }
    }
}

func Login(ctx *gin.Context, userParam *req.UserLoginOrRegisterParam) {
    switch userParam.LoginType {
    case 0:
        LoginByMobile(ctx, userParam)
    case 1:
        LoginByPassword(ctx, userParam)
    default:
        response.FailWithDetailed(gin.H{"reload": true}, "登录方式指定有错误", ctx)
        ctx.Abort()
        return
    }
}
func Register(ctx *gin.Context, userParam *req.UserLoginOrRegisterParam) {
    if err := dao.InsertUser(user.User{UserMobile: userParam.UserMobile}); err != nil {
        response.FailWithMessage(err.Error(), ctx)
        ctx.Abort()
        return
    }
    LoginByMobile(ctx, userParam)
}

func ConvUUIDToStr(uid uuid.UUID) string {
    bytes := uid.Bytes()
    return string(bytes)
}

func LoginByMobile(ctx *gin.Context, userParam *req.UserLoginOrRegisterParam) {
    // 生成token
    token, err := utils.CreateToken(userParam.UserMobile)
    if err != nil {
        response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
        ctx.Abort()
        return
    }
    ctx.SetCookie("token", token, 60*5, "", "", false, true)
    response.OkWithDetailed(gin.H{"token": token}, "登录成功", ctx)
    ctx.Next()
    return
}

func LoginByPassword(ctx *gin.Context, userParam *req.UserLoginOrRegisterParam) {
    // 生成token
    if userParam.Password != "123" {
        response.FailWithDetailed(gin.H{"reload": true}, "密码错误", ctx)
        ctx.Abort()
        return
    }
    token, err := utils.CreateToken(userParam.UserMobile)
    if err != nil {
        response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
        ctx.Abort()
        return
    }
    response.OkWithDetailed(gin.H{"token": token}, "登录成功", ctx)
    ctx.Next()
    return
}

func GetUserList() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        response.OkWithDetailed(response.PageResult{
            List:  nil,
            Total: 10,
        }, "获取成功", ctx)
    }
}
