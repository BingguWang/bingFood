package main

import (
    "log"

    "bingFood/initialize"
    "bingFood/router"

    "github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
    r = router.SetupRouter()
    initialize.Viper()
    initialize.MySql()
    initialize.Redis()
    initialize.Nsq()
}
func main() {
    if err := r.Run(":8088"); err != nil {
        log.Fatal(err)

    }
}
