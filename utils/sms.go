package utils

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/pkg/errors"
    "log"
    "math/rand"
    "orderModule/global"
    "orderModule/initialize"
    "strconv"
    "time"
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
    if err := InsertToRedis(code, mobileNumber); err != nil {
        log.Printf("insertToRedis failed, err : %v", err)
        return "", err
    }
    buildContent(code,mobileNumber)
    return code, nil // 这里直接返回就假装是用户收到短信了
}

func buildContent(code,mobile string) string {
    return "【bingFood】尊敬的"+mobile+"用户，您的验证码为：" + code + "，该验证码 " +strconv.Itoa(MsgValidTime) + " 分钟内有效，请勿泄漏于他人。"
}

func InsertToRedis(code, mobileNumber string) error {
    ctx := context.TODO()
    key := MsgPrefix + "-" + mobileNumber
    value := code
    cli := redis.NewClient(&redis.Options{
        Addr:     initialize.Addr,
        Password: "wb430481", // no password set
        DB:       0,          // 0 means to use default DB
    })
    if _, err := cli.Set(ctx, key, value, time.Minute*5).Result(); err != nil {
        log.Printf("set k-v failed, key : %v , value : %v \n", key, value)
        return err
    }
    return nil
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
