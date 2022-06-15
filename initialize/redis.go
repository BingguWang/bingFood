package initialize

import (
    "bingFood/global"
    "context"
    "github.com/go-redis/redis/v8"
    "log"
)

const Addr = "1.14.163.5:6379"

func Redis() {
    redisConf := global.GVA_CONFIG.Redis
    client := redis.NewClient(&redis.Options{
        Addr:     redisConf.Addr,
        Password: redisConf.Password, // no password set
        DB:       redisConf.DB,       // 0 means to use default DB
    })
    pong, err := client.Ping(context.Background()).Result()
    if err != nil {
        log.Printf("redis connect ping failed, err: %v", err)
    } else {
        log.Printf("redis connect ping response:%v \"pong\"", pong)
        global.GVA_REDIS = client
    }
}
