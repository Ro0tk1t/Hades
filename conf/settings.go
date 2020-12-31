package conf

import (
    "hades/models"

    "strings"
    "time"

    log "github.com/sirupsen/logrus"
    nested "github.com/antonfisher/nested-logrus-formatter"
)


var (
    Debug = false
    Timeout = 10 * time.Second
    Threads = 20
    CheckNone = false
    Quit = false

    IPs[]models.Target
    Users[]string
    Pass[]string
    Output string
)

var (
    PortService = map[int]string{
        21:    "FTP",
        22:    "SSH",
        23:    "TELNET",
        80:    "HTTPBASIC",
        161:   "SNMP",
        445:   "SMB",
    //    1433:  "MSSQL",
    //    3306:  "MYSQL",
    //    5432:  "POSTGRESQL",
        6379:  "REDIS",
    //    9200:  "ELASTICSEARCH",
    //    27017: "MONGODB",
    }
    OnlyNeedPass = map[string]bool{
        "SNMP": true,
        "REDIS": true,
    }
    SupportProto map[string]bool
    UDPProtocol = map[string]bool{
        "SNMP": true,
    }

    AliveTarget []models.Target
    SuccedResults []models.Result
)

func init(){
    log.SetFormatter(&nested.Formatter{
        HideKeys:    true,
        FieldsOrder: []string{"component", "category"},
    })

    SupportProto = make(map[string]bool)
    for _, proto := range PortService{
        SupportProto[strings.ToUpper(proto)] = true
    }
}
