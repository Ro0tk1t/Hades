package plugins

import (
    "fmt"
    "net"
    "bytes"
    "errors"
    "hades/models"
)

func BruteFTP(target models.Target, user, pass string)(result models.Result, err error){
    err = errors.New("incorrect password")
    conn, err_ := net.Dial("tcp", fmt.Sprintf("%v:%v", target.IP, target.Port))
    if err_ != nil{
        return
    }
    defer conn.Close()
    var (
        n int
        buf[512]byte
        fflag = []byte("220 ")
        pflag = []byte("331 ")
        ok = []byte("230 ")
    )

    n, err_ = conn.Read(buf[0:])
    if err_ == nil && bytes.HasPrefix(buf[0:n], fflag){
        _, err_ = conn.Write([]byte(fmt.Sprintf("USER %v\r\n", user)))
        if err_ == nil{
            n, err_ = conn.Read(buf[0:])
            if err_ == nil{
                if bytes.HasPrefix(buf[0:n], pflag){
                    _, err_ = conn.Write([]byte(fmt.Sprintf("PASS %v\r\n", pass)))
                    if err_ == nil{
                        n, err_ = conn.Read(buf[0:])
                        if err_ == nil{
                            if bytes.HasPrefix(buf[0:n], ok){
                                result.User = user
                                result.Pass = pass
                                result.Target = target
                                err = nil
                                return
                            }
                        }
                    }
                }
            }
        }
    }

    return
}
