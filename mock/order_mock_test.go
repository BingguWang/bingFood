package mock

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "testing"
)

/**
  模拟订单相关请求
*/
import (
    "fmt"
)

// curl -H "Content-Type:application/json" -XGET http://127.0.0.1:8088/login
func TestOrder(t *testing.T) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", HostPort+"/login", nil)
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

// curl -H "Content-Type:application/json" -XGET http://127.0.0.1:8088/orders/123/wang/bing
func TestOrder2(t *testing.T) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", HostPort+"/orders/123/wang/bing", nil)
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

// curl -X POST -H "Content-Type:application/x-www-form-urlencoded" -d 'userid=1&userName=wd' http://127.0.0.1:8088/order/list
func TestOrder3(t *testing.T) {
    client := &http.Client{}
    var data = strings.NewReader(`userid=1&userName=wd`)
    req, err := http.NewRequest("POST", HostPort+"/order/list", data)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

// curl -X POST -H "Content-Type:application/json" -d '{"userid":1,"userName":"wd"}' http://127.0.0.1:8088/order/lists
func TestOrder4(t *testing.T) {
    client := &http.Client{}
    var data = strings.NewReader(`{"userid":"1" ,"userName":"wbing"}`)
    req, err := http.NewRequest("POST", HostPort+"/user/add", data)
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

func jsonPost(route, postData string) string {
    client := &http.Client{}
    var data = strings.NewReader(postData)
    req, err := http.NewRequest("POST", route, data)
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
    return string(bodyText)
}

func urlGet(route, getData string) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", HostPort+"/orders/123/wang/bing", nil)
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
