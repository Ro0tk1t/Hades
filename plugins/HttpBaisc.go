package plugins

import (
    "fmt"
    "time"
    "errors"
    "strings"
    "net/http"
    "math/rand"
    "hades/models"
)

func BruteHttpBasic(target models.Target, user, pass string)(result models.Result, err error){
    uas := []string{
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.56 Safari/537.36",
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
        "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:40.0) Gecko/20100101 Firefox/40.0",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:41.0) Gecko/20100101 Firefox/41.0",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11) AppleWebKit/601.1.27 (KHTML, like Gecko) Chrome/47.0.2526.106 Safari/601.1.27",
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Ubuntu/11.10 Chromium/27.0.1453.93 Chrome/27.0.1453.93 Safari/537.36",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11) AppleWebKit/601.1.27 (KHTML, like Gecko) Version/8.1 Safari/601.1.27",
        "Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2376.0 Safari/537.36 OPR/31.0.1857.0",
    }
    rand.Seed(time.Now().UnixNano())
    
    var protocol, url string
    if target.URL != ""{
        url = target.URL
    } else {
        if target.Protocol == ""{
            protocol = "http"
        } else {
            protocol = target.Protocol
        }

        if strings.Contains(strings.ToLower(protocol), "https"){
            protocol = "https"
        } else if strings.Contains(strings.ToLower(protocol), "http") {
            protocol = "http"
        }
        url = fmt.Sprintf("%v://%v:%v", protocol, target.IP, target.Port)
    }
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("User-Agent", uas[rand.Intn(len(uas))])
    req.SetBasicAuth(user, pass)
    res, err := client.Do(req)
    if err == nil && res.StatusCode == 200{
        result.User = user
        result.Pass = pass
        result.Target = target
    } else {
        err = errors.New("incorrect password")
    }
    defer res.Body.Close()
    return
}
