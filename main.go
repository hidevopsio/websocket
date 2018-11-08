
//go:generate go build -o websocket

package main

import (
	"hidevops.io/hiboot/pkg/app/cli"
	"hidevops.io/websocket/cmd"
)

func main()  {
	cli.NewApplication(cmd.NewWebsocketCommand).
		Run()
}