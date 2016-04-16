// +build linux darwin

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"runtime"
	"sort"
)

var cmds = []string{
	"login", "logout", "cd", "pwd", "get", "put",
	"ls", "rm", "switch", "info", "mkdir", "services",
}

func main() {
	app := cli.NewApp()
	app.Name = "upx"
	app.Usage = "a tool for managing files in UPYUN"
	app.Author = "Hongbo.Mo"
	app.Email = "zjutpolym@gmail.com"
	app.Version = fmt.Sprintf("v0.1.1 %s/%s", runtime.GOOS, runtime.GOARCH)
	app.Commands = make([]cli.Command, 0)

	sort.Strings(cmds)
	for _, cmd := range cmds {
		cm, exist := CmdMap[cmd]
		if exist {
			Cmd := cli.Command{
				Name:  cmd,
				Usage: cm.Desc,
				Action: func(c *cli.Context) {
					if c.Command.FullName() != "login" && driver == nil {
						fmt.Println("Log in first.")
						os.Exit(-1)
					}
					cm.Func(c.Args())
				},
			}
			if cm.Alias != "" {
				Cmd.Aliases = []string{cm.Alias}
			}
			app.Commands = append(app.Commands, Cmd)
		}
	}

	app.Run(os.Args)
}
