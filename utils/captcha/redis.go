package captcha

import (
    "bingFood/global"
    "bingFood/utils"
    "context"
    "github.com/mojocn/base64Captcha"
    "log"
    "time"
)
/**
    用于人机校验
 */
func NewDefaultRedisStore() *RedisStore {
    return &RedisStore{
        Expiration: time.Second * 180,
        PreKey:     "CAPTCHA_",
    }
}

type RedisStore struct {
    Expiration time.Duration
    PreKey     string
    Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
    rs.Context = ctx
    return rs
}

func (rs *RedisStore) Set(id string, value string) error {
    if err := utils.InsertToRedis(rs.Context, rs.PreKey+id, value, rs.Expiration); err != nil {
        log.Printf("RedisStoreSetError! %v", err.Error())
        return err
    }
    return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
    val, err := global.GVA_REDIS.Get(rs.Context, key).Result()
    if err != nil {
        log.Printf("RedisStoreGetError! %v", err.Error())
        return ""
    }
    if clear {
        err := global.GVA_REDIS.Del(rs.Context, key).Err()
        if err != nil {
            log.Printf("RedisStoreClearError! %v", err.Error())
            return ""
        }
    }
    return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
    key := rs.PreKey + id
    v := rs.Get(key, clear)
    return v == answer
}
