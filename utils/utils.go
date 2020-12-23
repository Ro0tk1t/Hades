package utils

import (
    "hades/models"
    "hades/conf"
    "encoding/json"
    "net/url"
    "strconv"
    "strings"
    "bufio"
    "sync"
    "fmt"
    "net"
    "log"
    "os"
)

var m sync.Mutex


func SplitIP(line string){
    ipPort := strings.TrimSpace(line)
    URL, err := url.Parse(ipPort)
    if err != nil || URL.Hostname() == "" || URL.Port() == ""{
        t := strings.Split(ipPort, ":")
        ip := t[0]
        portProtocol := strings.Split(t[1], "|")
        if len(portProtocol) == 2{
            port, _ := strconv.Atoi(portProtocol[0])
            protocol := strings.ToUpper(portProtocol[1])
            if conf.SupportProto[protocol]{
                target := models.Target{IP: ip, Port: port, Protocol: protocol}
                conf.IPs = append(conf.IPs, target)
            } else {
                log.Printf("Not Support %v, ignore: %v:%v", protocol, ip, port)
            }
        } else {
            port, err := strconv.Atoi(portProtocol[0])
            if err == nil{
                protocol, ok := conf.PortService[port]
                if ok && conf.SupportProto[protocol]{
                    target := models.Target{IP: ip, Port: port, Protocol: protocol}
                    conf.IPs = append(conf.IPs, target)
                }
            }
        }
        return
    }

    // parse url format
    ip := URL.Hostname()
    protocol := strings.ToUpper(URL.Scheme)
    if strings.Contains(strings.ToLower(protocol), "https"){
        URL.Scheme = "https"
    } else if strings.Contains(strings.ToLower(protocol), "http") {
        URL.Scheme = "http"
    }
    port, _ := strconv.Atoi(URL.Port())
    if conf.SupportProto[protocol]{
        target := models.Target{IP: ip, Port: port, Protocol: protocol, URL: fmt.Sprintf("%v", URL)}
        conf.IPs = append(conf.IPs, target)
    } else {
        log.Printf("Not Support %v, ignore: %v:%v", protocol, ip, port)
    }

}


func GetIPs(fn string){
    fd, err := os.Open(fn)
    if err != nil{
        log.Println(err)
        return
    }
    defer fd.Close()

    scanner := bufio.NewScanner(fd)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        line := scanner.Text()
        if line == ""{
            continue
        }
        SplitIP(line)
    }
}


func GetFileData(fn string) []string{
    var datas []string
    fd, err := os.Open(fn)
    if err != nil{
        log.Println(err)
        return datas
    }
    defer fd.Close()
    scanner := bufio.NewScanner(fd)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan(){
        data := strings.TrimSpace(scanner.Text())
        if data != ""{
            datas = append(datas, data)
        }
    }
    return datas
}


func CheckAlive(targets []models.Target){
    log.Printf("checking alived target")
    wg := sync.WaitGroup{}
    wg.Add(len(targets))
    for _, target := range targets{
        go func(target models.Target){
            defer wg.Done()
            Check(target)
        }(target)
    }
    wg.Wait()
}


func Check(target models.Target){
    alive := false
    if conf.UDPProtocol[target.Protocol]{
        conn, err := net.DialTimeout("udp", fmt.Sprintf("%v:%v", target.IP, target.Port), conf.Timeout)
        if err == nil{
            alive = true
            conn.Close()
        }
    } else {
        conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", target.IP, target.Port), conf.Timeout)
        if err == nil{
            alive = true
            conn.Close()
        }
    }
    if alive{
        m.Lock()
        conf.AliveTarget = append(conf.AliveTarget, target)
        log.Println("found alived: ", target)
        m.Unlock()
    }
}


func WriteResults(){
    results, err := json.MarshalIndent(conf.SuccedResults, "", "  ")
    if err != nil{
        log.Fatal(err)
    } else {
        if conf.Output == ""{
            fmt.Println("")
            log.Println("Results is:")
            fmt.Println(string(results))
        } else {
            fd, err := os.OpenFile(conf.Output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
            defer fd.Close()
            if err != nil{
                log.Fatal(err)
            } else {
                _, err := fd.Write(results)
                if err != nil{
                    log.Fatal(err)
                }
            }
        }
    }
}
