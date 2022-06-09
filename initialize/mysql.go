package initialize

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "orderModule/global"
)

func MySql() {
    db, err := sql.Open("mysql", "root:1234@tcp(1.14.163.5:3306)/bingFood")
    if err != nil {
        log.Printf("open mysql failed, err: %v", err)
    } else {
        log.Printf("open mysql success")
        global.MYSQL_DB = db
    }
}
