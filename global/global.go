package global

import (
    "database/sql"
    "github.com/go-redis/redis/v8"
    _ "github.com/go-sql-driver/mysql"
)

var (
    GVA_REDIS *redis.Client
    MYSQL_DB  *sql.DB
    //GVA_DB     *gorm.DB
    //GVA_DBList map[string]*gorm.DB
    //GVA_CONFIG config.Server
    //GVA_VP     *viper.Viper
    //// GVA_LOG    *oplogging.Logger
    //GVA_LOG                 *zap.Logger
    //GVA_Timer               timer.Timer = timer.NewTimerTask()
    //GVA_Concurrency_Control             = &singleflight.Group{}
    //
    //BlackCache local_cache.Cache
    //lock       sync.RWMutex
)
