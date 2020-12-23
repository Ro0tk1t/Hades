package plugins

import (
    "hades/models"
)


type ScanFunc func(target models.Target, user, pass string)(result models.Result, err error)
var ScanFuncMap map[string]ScanFunc

func init(){
    ScanFuncMap = make(map[string]ScanFunc)
    ScanFuncMap["FTP"] = BruteFTP
    ScanFuncMap["SSH"] = BruteSSH
    ScanFuncMap["TELNET"] = BruteTelnet
    ScanFuncMap["HTTPBASIC"] = BruteHttpBasic
    ScanFuncMap["REDIS"] = BruteRedis
    ScanFuncMap["SMB"] = BruteSMB
    ScanFuncMap["SNMP"] = BruteSNMP
    // TODO: plugin
}
