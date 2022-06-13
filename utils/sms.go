package utils

import (
    "bingFood/global"
    "github.com/pkg/errors"
    "time"

    "context"
    "fmt"
    "log"
    "math/rand"
    "strconv"
)

const (
    MsgValidTime = 5
    MsgPrefix    = "sms:"
)

func SendMsg(mobileNumber string) (string, error) {
    // 直接自动生成随机验证码，模拟发送验证码SDK
    code := fmt.Sprintf("%05v", rand.Int31n(99999))
    // 如果要防止出现重复，可以放到一个map里，每次先查询map是否存在了，然后每5分钟的时候清除一次map，可以有效减少重复

    // 存入redis
    key := MsgPrefix + "-" + mobileNumber
    if err := InsertToRedis(context.TODO(), key, code, time.Minute*5); err != nil {
        log.Printf("insertToRedis failed, err : %v", err)
        return "", err
    }
    buildContent(code, mobileNumber)
    return code, nil // 这里直接返回就假装是用户收到短信了
}

func buildContent(code, mobile string) string {
    return "【bingFood】尊敬的" + mobile + "用户，您的验证码为：" + code + "，该验证码 " + strconv.Itoa(MsgValidTime) + " 分钟内有效，请勿泄漏于他人。"
}

func CheckValidCode(code, userMobile string) (bool, error) {
    if code == "" {
        return false, errors.New("验证码为空")
    }
    key := MsgPrefix + "-" + userMobile
    value := code
    res, err := global.GVA_REDIS.Get(context.TODO(), key).Result()
    if err != nil {
        log.Printf("set k-v failed, key : %v , value : %v \n", key, value)
        return false, err
    }
    if value == res {
        return true, nil
    }
    return false, errors.New("验证码错误")
}
