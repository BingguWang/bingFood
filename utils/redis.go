package utils

import (
    "bingFood/global"
    "context"
    "log"
    "time"
)

func InsertToRedis(ctx context.Context, key, value string, expire time.Duration) error {
    cli := global.GVA_REDIS
    if _, err := cli.Set(ctx, key, value, expire).Result(); err != nil {
        log.Printf("set k-v failed, key : %v , value : %v \n", key, value)
        return err
    }
    return nil
}
