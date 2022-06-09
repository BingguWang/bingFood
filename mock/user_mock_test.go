package mock

import (
    "fmt"
    "sync"
    "testing"
)

// curl -X POST -H "Content-Type:application/json" -d '{"userMobile":"15759216850"}' http://127.0.0.1:8088/user/getCode
const (
    getCodeRoute    = "http://127.0.0.1:8088/user/getCode"
    getCodePostData = `{"userMobile":"1000"}`
)

func Test_GetValidCode(t *testing.T) {
    var wg sync.WaitGroup
    mp := sync.Map{}
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        i := i
        go func() {
            defer wg.Done()
            getCodePostData := fmt.Sprintf(`{"userMobile":"%04v"}`, i)
            resp := jsonPost(getCodeRoute, getCodePostData)
            mp.Store(resp, struct{}{})
            fmt.Println(resp)
        }()
    }
    wg.Wait()
    count := 0
    mp.Range(func(k, v interface{}) bool {
        count++
        return true
    })
    fmt.Println(count) // 可以看到这种自己模拟发验证码并发时会有出现重复的可能性，直接用服务商的SDK比较好
}

