package plugins

import (
    "hades/models"
    "errors"
    "strings"
    "fmt"
    "net"
)

// https://blog.csdn.net/pssmart/article/details/51500422
func BruteTelnet(target models.Target, user, pass string)(result models.Result, err error){
    conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", target.IP, target.Port))
    if err == nil{
        defer conn.Close()
        var buf [2048]byte
        var n int
        n, err = conn.Read(buf[0:])

        buf[1] = 252
        buf[4] = 252
        buf[7] = 252
        buf[10] = 252

        n, err = conn.Write(buf[0:n])
        if err == nil{
            n, err = conn.Read(buf[0:])
            if err == nil{
                buf[1] = 252
                buf[4] = 251
                buf[7] = 252
                buf[10] = 254
                buf[13] = 252
                n, err = conn.Write(buf[0:n])
                if err == nil{
                    n, err = conn.Read(buf[0:])
                    if err == nil{
                        buf[1] = 252
                        buf[4] = 252
                        n, err = conn.Write(buf[0:n])
                        if err == nil{
                            n, err = conn.Read(buf[0:])
                            if err == nil{
                                n, err = conn.Write([]byte(user + "\n"))
                                if err == nil{
                                    n, err = conn.Read(buf[0:])
                                    if err == nil{
                                        conn.Write([]byte(pass + "\n"))
                                        if err == nil{
                                            n, err = conn.Read(buf[0:])
                                            if err == nil{
                                                n, err = conn.Read(buf[0:])
                                                if strings.Contains(string(buf[0:n]), "Last login"){
                                                    result.User = user
                                                    result.Pass = pass
                                                    result.Target = target
                                                } else {
                                                    err = errors.New("incorrect password")
                                                    return
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    return
}

