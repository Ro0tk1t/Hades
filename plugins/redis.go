package plugins

import (
    "fmt"
    "net"
    "errors"
    "strings"
    "hades/models"
)

func BruteRedis(target models.Target, user, pass string)(result models.Result, err error){
    conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", target.IP, target.Port))
    if err != nil{
        return
    }

    var (
        cmd []byte
        buf [1024]byte
        flag_p = "+PONG"
        flag_ok = "+OK"
    )
    if pass != ""{
        cmd = []byte(fmt.Sprintf("auth %v\r\n", pass))
        _, err = conn.Write(cmd[0:])
        if err != nil{
            return
        }

        _, err = conn.Read(buf[0:])
        if err != nil{
            return
        }
        if strings.HasPrefix(string(buf[0:]), flag_ok){
            result.Pass = pass
            result.Target = target
            return
        }
        err = errors.New("incrrect password")
    } else {
        cmd = []byte("ping\r\n")
        _, err = conn.Write(cmd[0:])
        if err != nil{
            return
        }

        _, err = conn.Read(buf[0:])
        if err != nil{
            return
        }
        if strings.HasPrefix(string(buf[0:]), flag_p){
            result.Pass = pass
            result.Target = target
            return
        }
        err = errors.New("incrrect password")
    }
    return
}

