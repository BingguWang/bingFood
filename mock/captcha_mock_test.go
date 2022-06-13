package mock

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "testing"
)

// curl -X POST -H "Content-Type:application/json" -d '{}' http://127.0.0.1:8088/user/getCaptcha
func Test_Captcha(t *testing.T) {
    client := &http.Client{}
    var data = strings.NewReader(`{}`)
    req, err := http.NewRequest("POST", "http://127.0.0.1:8088/user/getCaptcha", data)
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
