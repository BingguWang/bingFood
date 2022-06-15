package global

import (
    "bingFood/config"
    "github.com/go-redis/redis/v8"
    _ "github.com/go-sql-driver/mysql"
    "github.com/spf13/viper"
    "gorm.io/gorm"
)

var (
    GVA_REDIS  *redis.Client
    MYSQL_DB   *gorm.DB
    GVA_VP     *viper.Viper
    GVA_CONFIG config.ServerConfig
    //GVA_DB     *gorm.DB
    //GVA_DBList map[string]*gorm.DB
    //// GVA_LOG    *oplogging.Logger
    //GVA_LOG                 *zap.Logger
    //GVA_Timer               timer.Timer = timer.NewTimerTask()
    //GVA_Concurrency_Control             = &singleflight.Group{}
    //
    //BlackCache local_cache.Cache
    //lock       sync.RWMutex
)
