package initialize

import (
    "context"
    "github.com/go-redis/redis/v8"
    "log"
    "orderModule/global"
)

const Addr = "1.14.163.5:6379"

func Redis() {
    // TODO redis的这些配置参数应该用yaml管理
    client := redis.NewClient(&redis.Options{
        Addr:     Addr,
        Password: "wb430481", // no password set
        DB:       0,          // 0 means to use default DB
    })
    pong, err := client.Ping(context.Background()).Result()
    if err != nil {
        log.Printf("redis connect ping failed, err: %v", err)
    } else {
        log.Printf("redis connect ping response:%v \"pong\"", pong)
        global.GVA_REDIS = client
    }
}
