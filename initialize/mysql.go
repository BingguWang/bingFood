package initialize

import (
    "bingFood/global"
    _ "github.com/go-sql-driver/mysql"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"
)

//TODO 要用viper和yaml来管理配置
func MySql() {
    defer func() {
        if e := recover(); e != nil {
            log.Printf("open mysql failed, err: %v", e)
        }
    }()
    dsn := "root:1234@tcp(1.14.163.5:3306)/bingFood?charset=utf8mb4&parseTime=True&loc=Local"

    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second,   // 慢 SQL 阈值
            LogLevel:      logger.Silent, // Log level
            Colorful:      true,          // 彩色打印
        },
    )

    db, err := gorm.Open(mysql.New(mysql.Config{
        DSN:                       dsn,  // DSN data source name
        DefaultStringSize:         191,  // string 类型字段的默认长度
        SkipInitializeWithVersion: true, // 根据版本自动配置
    }), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,
        QueryFields: true, // 这样查询的时候是用字段名称而不是*
        Logger:      newLogger.LogMode(logger.Info), // 开启info级别的日志
    })

    if err != nil {
        panic(err)
    }
    log.Printf("open mysql success")

    global.MYSQL_DB = db

}
