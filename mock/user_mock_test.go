package mock

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "sync"
    "testing"
)

// curl -X POST -H "Content-Type:application/json" -d '{"userMobile":"15759216850"}' http://127.0.0.1:8088/user/getCode
const (
    getCodeRoute    = HostPort + "/user/getCode"
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

// curl -X POST -H "Content-Type:application/json" -d '{}' http://127.0.0.1:8088/userList
func Test_UserList(t *testing.T) {
    client := &http.Client{}
    var data = strings.NewReader(`{"userid":"1" ,"userName":"wbing"}`)
    req, err := http.NewRequest("POST", HostPort+"/user/api/userList", data)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."+
        "eyJVc2VyTW9iaWxlIjoiMTU3NTkyMTY4NTAiLCJleHAiOjE2NTQ4NDY2NTUsImlhdCI6MTY1NDg0NjM1NSwiaXNzIjoiMTI3LjAuMC4xIiwic3ViIjoidXNlciB0b2tlbiJ9."+
        "MAJUtpU2vWx5vFl7JeENJMKHSfiqx4514vAbT_PJUdo")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    bodyText, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", bodyText)
}

const HostPort = "http://127.0.0.1:8088"

// curl -X POST -H "Content-Type:application/json" -d '{}' http://127.0.0.1:8088/user/login
func Test_UserLogin(t *testing.T) {
    client := &http.Client{}
    var data = strings.NewReader(`{"userMobile":"15759216850","isLogin":1}`)
    req, err := http.NewRequest("POST", HostPort+"/user/login", data)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    bodyText, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", bodyText)
}

