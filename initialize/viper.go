package initialize

import (
    "bingFood/global"
    "bingFood/utils"
    "flag"
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "os"
)

const (
    ConfigEnv  = "GVA_CONFIG"
    ConfigFile = "config.yaml"
)

func Viper(path ...string) *viper.Viper {
    var config string
    if len(path) == 0 {
        flag.StringVar(&config, "c", "", "choose config file.")
        flag.Parse()
        if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
            if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
                config = ConfigFile
                fmt.Printf("您正在使用config的默认值,config的路径为%v\n", ConfigFile)
            } else {
                config = configEnv
                fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
            }
        } else {
            fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
        }
    } else {
        config = path[0]
        fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
    }

    v := viper.New()
    v.SetConfigFile(config)                  // 设置配置文件路径
    v.SetConfigType("yaml")                  // 设置配置文件类型
    if err := v.ReadInConfig(); err != nil { // 根据设置的文件路径读取配置文件
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    v.OnConfigChange(func(e fsnotify.Event) { // 实时查看配置文件，变化就重新加载
        fmt.Println("config file changed:", e.Name)
        if err := v.Unmarshal(&global.GVA_CONFIG); err != nil { // 读取到配置后要解析到定义的配置实体里
            fmt.Println(err)
        }
    })
    v.WatchConfig() // 实时查看配置文件，变化就重新加载

    if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
        fmt.Println(err)
    }
    fmt.Println(utils.ToJsonString(global.GVA_CONFIG))

    return v
}
