package initialize

import (
    "bingFood/global"
    "github.com/bwmarrin/snowflake"
    "log"
)

func SnowFlakeNode() {
    var node *snowflake.Node
    node, err := snowflake.NewNode(1) // 新建一个节点号为1的node
    if err != nil {
        log.Fatalf("redis connect ping failed, err: %v", err)
    } else {
        global.SNOW_NODE = node
    }

}
