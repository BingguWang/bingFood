package middleware

import (
    "bingFood/common/response"
    "bingFood/utils"
    "bingFood/utils/captcha"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/mojocn/base64Captcha"
    "log"
)

// 人机校验验证码
var store = captcha.NewDefaultRedisStore()

//var store = base64Captcha.DefaultMemStore

func GetCaptcha() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        param := CaptchaReq{}
        ctx.BindJSON(&param)

        var driver base64Captcha.Driver

        // 可以根据不同的需要创建各种类型的captcha
        driver = base64Captcha.NewDriverDigit(80, 230, 4, 0.7, 100)
        //c := base64Captcha.NewCaptcha(driver, store)
        c := base64Captcha.NewCaptcha(driver, store.UseWithCtx(ctx))
        id, b64s, err := c.Generate()
        if err != nil {
            log.Printf("验证码获取失败! %v", err.Error())
            response.FailWithMessage("验证码获取失败", ctx)
        }
        response.OkWithDetailed(CaptchaResponse{
            CaptchaId:     id,
            PicPath:       b64s,
            CaptchaLength: 5,
        }, "验证码获取成功", ctx)
    }

}

// TODO 其实并不需要把验证captcha单独作为一个middleware，其实应该把verify作为service里，这里只是为了体现效果
func VerifyCaptcha() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var param CaptchaReq
        ctx.BindJSON(&param)
        store.UseWithCtx(ctx)
        fmt.Println(utils.ToJsonString(param))
        if !store.Verify(param.Id, param.VerifyValue, true) {
            response.FailWithMessage("图形验证码填写错误", ctx)
            ctx.Abort()
            return
        }
        response.OkWithMessage("图形验证码填写正确", ctx)
    }
}

//configJsonBody json request body.
type CaptchaReq struct {
    Id          string // 生成的captcha的id
    CaptchaType string
    VerifyValue string // 用户填写的captcha值
}
type CaptchaResponse struct {
    CaptchaId     string `json:"captchaId"`
    PicPath       string `json:"picPath"`
    CaptchaLength int    `json:"captchaLength"`
}
