package plugins

import (
    "golang.org/x/crypto/ssh"
    "hades/models"
    "hades/conf"
    "fmt"
    "net"
)

func BruteSSH(target models.Target, user, pass string)(result models.Result, err error){
    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.Password(pass),
        },
        Timeout: conf.Timeout,
        HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey)error{
            return nil
        },
    }

    client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", target.IP, target.Port), config)
    if err == nil{
        defer client.Close()
        session, err := client.NewSession()
        errRet := session.Run("pwd")
        if err == nil && errRet == nil{
            defer session.Close()
            result.User = user
            result.Pass = pass
            result.Target = target
        }
    }
    return
}
