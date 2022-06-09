package service

import (
    "log"
    "orderModule/entity/user"
    "orderModule/global"
    "time"
)

func InsertUser(user user.User) (int64, error) {
    db := global.MYSQL_DB
    tx, err := db.Begin()
    if err != nil {
        panic("open tx failed," + err.Error())
    }
    result, err := tx.Exec("insert into t_user(user_id, user_wx_number)values(?,?)", user.UserId, user.UserWxNum)
    if err != nil {
        log.Printf("%v", err)
        tx.Rollback()
        return -1, err
    }
    time.Sleep(1 * time.Second)
    tx.Commit()

    return result.LastInsertId()
}
