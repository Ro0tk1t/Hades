package plugins

import (
    "fmt"
    "net"
    "bytes"
    "hades/models"
)

// https://www.cnblogs.com/chloneda/p/snmp-protocol.html
func BruteSNMP(target models.Target, user, pass string)(result models.Result, err error){
    conn, err := net.Dial("udp", fmt.Sprintf("%v:%v", target.IP, target.Port))
    if err != nil{
        return
    }
    var (
        buf [512]byte
        flag = []byte{0x2b,0x06,0x01,0x02,0x01,0x01,0x05,0x00}
    )
    length := len(pass)
    pkg_len := 0x23 + length
    get := []byte{0x30, uint8(pkg_len), 0x02, 0x01, 0x01, 0x04, uint8(length)}
    get = append(get, []byte(pass)...)
    get = append(get, []byte{0xa0,0x1c,0x02,0x04,0x70,0x8c,0x32,0x17,0x02,0x01,0x00,0x02,0x01,0x00,0x30,0x0e,0x30,0x0c,0x06,0x08}...)
    get = append(get, flag...)
    get = append(get, []byte{0x05,0x00}...)
    _, err = conn.Write(get[0:])
    if err != nil{
        return
    }
    n, err := conn.Read(buf[0:])
    if  err != nil{
        return
    }
    hostname := string(buf[44:n])

    flag = append(flag, []byte{0x04, uint8(len(hostname))}...)
    if bytes.Contains(buf[0:n], flag) && bytes.Contains(buf[0:n], []byte(pass)){
        result.Pass = pass
        result.Target = target
        return
    }
    return
}
