package main

import (
	"hidevops.io/hiboot/pkg/app/cli"
	"hidevops.io/websocket-cli/cmd"
)

func main()  {
	cli.NewApplication(cmd.NewWebsocketCommand).
		Run()
}