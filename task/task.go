package task

import (
    "time"
    "sync"
    "context"
    "strings"

    "hades/conf"
    "hades/models"
    "hades/plugins"
    . "hades/utils"

    "github.com/urfave/cli/v2"
    log "github.com/sirupsen/logrus"
)

var m = sync.Mutex{}

func Scan(ctx *cli.Context)(err error){
    if ctx.IsSet("debug"){
        conf.Debug = ctx.Bool("debug")
    }
    if ctx.IsSet("timeout"){
        conf.Timeout = time.Duration(ctx.Int("timeout")) * time.Second
    }
    if ctx.IsSet("threads"){
        conf.Threads = ctx.Int("threads")
    }
    if ctx.IsSet("quit"){
        conf.Quit = ctx.Bool("quit")
    }
    if ctx.IsSet("ip"){
        SplitIP(ctx.String("ip"))
    }
    if ctx.IsSet("ipfile"){
        GetIPs(ctx.String("ipfile"))
    }
    if ctx.IsSet("none"){
        conf.Pass = append(conf.Users, "")
    }
    if ctx.IsSet("output"){
        conf.Output = ctx.String("output")
    }
    conf.Users = append(conf.Users, GetFileData(ctx.String("user_file"))...)
    conf.Pass = append(conf.Pass, GetFileData(ctx.String("pass_file"))...)

    if conf.Users != nil && conf.Pass != nil{
        CheckAlive(conf.IPs)
        CheckPass()
    }
    return nil
}


func CheckPass(){
    wgi := sync.WaitGroup{}
    wgi.Add(len(conf.AliveTarget))
    for _, target := range conf.AliveTarget{
        go Brute(target, &wgi)
        wgi.Wait()
    }

    WriteResults()
}


func Brute(target models.Target, wgi *sync.WaitGroup){
    log.Info("start 2 brute force  ", target)
    wgt := sync.WaitGroup{}
    ch := make(chan struct{}, conf.Threads)
    right := false
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    if conf.OnlyNeedPass[target.Protocol]{
        wgt.Add(len(conf.Pass))
        for _, pass := range conf.Pass{
            select {
                case <- ctx.Done():
                    log.Info("stop brute force  ", target)
                    wgt.Done()
                default:
                    ch <- struct{}{}
                    go BruteForce(ctx, cancel, &wgt, ch, target, "", pass, &right)
            }
        }
    } else {
        wgt.Add(len(conf.Users) * len(conf.Pass))
        for _, user := range conf.Users{
            for _, pass := range conf.Pass{
                select {
                    case <- ctx.Done():
                        log.Info("stop brute force  ", target)
                        wgt.Done()
                    default:
                        ch <- struct{}{}
                        go BruteForce(ctx, cancel, &wgt, ch, target, user, pass, &right)
                }
            }
        }
    }
    wgt.Wait()
    wgi.Done()
}


func BruteForce(ctx context.Context, cancel context.CancelFunc, wgt *sync.WaitGroup, ch chan struct{}, target models.Target, user, pass string, right *bool){
    if conf.Quit && *right{
        wgt.Done()
        <- ch
        return
    } else {
        protocol := strings.ToUpper(target.Protocol)
        ctxB, _ := context.WithTimeout(ctx, conf.Timeout)
        ch0 := make(chan struct{}, 0)
        go func(){
            fn := plugins.ScanFuncMap[protocol]
            result, err := fn(target, user, pass)
            if err == nil{
                m.Lock()
                conf.SuccedResults = append(conf.SuccedResults, result)
                m.Unlock()
                if conf.Quit{
                    defer cancel()
                }
                log.Warn("found user: %v, pass: %v @ %v", user, pass, target)
                *right = true
            }
            ch0 <- struct{}{}
        }()
        select {
        case <- ch0:
            wgt.Done()
            <- ch
        case <- ctxB.Done():
            wgt.Done()
            <- ch
        }
    }
}
