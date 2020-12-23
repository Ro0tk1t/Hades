package plugins

import (
    "fmt"
    "net"
    "hades/models"
    "github.com/hirochachacha/go-smb2"
)


func BruteSMB(target models.Target, user, pass string)(result models.Result, err error){

    conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", target.IP, target.Port))
	if err != nil {
        return
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     user,
			Password: pass,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
        return
	}
	defer s.Logoff()

    result.User = user
    result.Pass = pass
    result.Target = target

    return
}
