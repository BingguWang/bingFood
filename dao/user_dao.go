package dao

import (
    user2 "bingFood/entity/user"
    "bingFood/global"
    "bingFood/utils"
    "log"

)

func GetUserByCond(user user2.User) []user2.UserEntity {
    log.Printf("exec GetUserByCond , args : %v", utils.ToJsonString(user))

    db := global.MYSQL_DB
    var userEntitys []user2.UserEntity
    db.Where("user_mobile = ?", user.UserMobile).Find(&userEntitys)

    return userEntitys
}

func InsertUser(user user2.User) error {
    log.Printf("exec InsertUser , args : %v", utils.ToJsonString(user.UserId))
    db := global.MYSQL_DB

    if err := db.Select("user_mobile").Create(&user).Error; err != nil {
        log.Printf("insert into db failed : %v", err.Error())
        return err
    }
    return nil
}
