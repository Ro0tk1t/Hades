package main

import (
    "github.com/urfave/cli/v2"
    "hades/cmd"
    "os"
)

func main(){
    app := &cli.App{
        Name: "hades scanner",
        Usage: "hades --help",
        Commands: []*cli.Command{&cmd.Scan,},
        Flags: cmd.Flags,
    }

    app.Run(os.Args)
}

