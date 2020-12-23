package cmd

import (
    "github.com/urfave/cli/v2"
    "hades/task"
)


var Scan = cli.Command{
    Name: "scan",
    Usage: "start a scan",
    Action: task.Scan,

}

var Flags = []cli.Flag{
    &cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Usage: "debug mode"},
    &cli.IntFlag{Name: "timeout", Aliases: []string{"t"}, Value: 10, Usage: "connection timeout"},
    &cli.IntFlag{Name: "threads", Aliases: []string{"T"}, Value: 20, Usage: "threads"},
    &cli.BoolFlag{Name: "none", Aliases: []string{"n"}, Usage: "check none pass"},
    &cli.BoolFlag{Name: "quit", Aliases: []string{"q"}, Usage: "quit when found a pass"},
    &cli.StringFlag{Name: "ipfile", Aliases: []string{"i"}, Usage: "ip list file"},
    &cli.StringFlag{Name: "ip", Aliases: []string{"I"}, Usage: "ip"},
    &cli.StringFlag{Name: "user_file", Aliases: []string{"u"}, Value: "users.txt", Usage: "user dict"},
    &cli.StringFlag{Name: "pass_file", Aliases: []string{"p"}, Value: "pass.txt", Usage: "password file"},
    &cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "result output file"},
}
