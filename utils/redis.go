package utils

import (
    "bingFood/global"
    "context"
    "fmt"
    "log"
    "time"
)

func InsertToRedis(ctx context.Context, key, value string, expire time.Duration) error {
    cli := global.GVA_REDIS
    fmt.Println(cli == nil)
    if _, err := cli.Set(ctx, key, value, expire).Result(); err != nil {
        log.Printf("set k-v failed, key : %v , value : %v \n", key, value)
        return err
    }
    return nil
}
