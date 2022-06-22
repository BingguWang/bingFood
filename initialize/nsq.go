package initialize

import (
    "bingFood/global"
    "github.com/nsqio/go-nsq"
)

func Nsq()  {
    global.NSQ_CONFIG = nsq.NewConfig()
}