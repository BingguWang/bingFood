package middleware

import (
    "bingFood/common/response"
    "bingFood/utils/captcha"
    "github.com/gin-gonic/gin"
    "github.com/mojocn/base64Captcha"
    "log"
)

// 人机校验验证码
var store = captcha.NewDefaultRedisStore()
//var store = base64Captcha.DefaultMemStore

func GetCaptcha() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var param GetCaptchaReq
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
        response.OkWithDetailed(GetCaptchaResponse{
            CaptchaId:     id,
            PicPath:       b64s,
            CaptchaLength: 5,
        }, "验证码获取成功", ctx)
    }

}

//configJsonBody json request body.
type GetCaptchaReq struct {
    Id          string
    CaptchaType string
    VerifyValue string
}
type GetCaptchaResponse struct {
    CaptchaId     string `json:"captchaId"`
    PicPath       string `json:"picPath"`
    CaptchaLength int    `json:"captchaLength"`
}
