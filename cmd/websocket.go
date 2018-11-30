package cmd

import (
	"hidevops.io/hiboot/pkg/app/cli"
	"fmt"
	"os"
	"net/url"
	"path/filepath"
	"io"
	"os/user"
	"github.com/chzyer/readline"
	"hidevops.io/websocket/ws"
)

const version = "1.0.0"

type websocketCommand struct {
	cli.RootCommand


	url string
	token string
	version bool
}

func NewWebsocketCommand() *websocketCommand  {
	c := new(websocketCommand)
	c.Use = "ws"
	c.Short = "websocket"
	c.Long = "websocket command line client "
	options := c.PersistentFlags()
	options.StringVarP(&c.token, "token", "t", "", "token. e.g. --token=the-token-string")
	options.StringVarP(&c.url, "url", "u", "", "websocket origin. e.g. --token=the-token-string")
	options.BoolVarP(&c.version, "version", "v", false, "print version, e.g. websocket version or websocket -v websocket --version")
	return c
}


func (c *websocketCommand) OnVersion(args []string) bool {
	fmt.Printf("websocket v%s\n", version)
	return false
}

func (c *websocketCommand) Run(args []string) error {
	if c.version {
		c.OnVersion(args)
		os.Exit(0)
	}

	if len(args) < 1 {
		c.Help()
		os.Exit(1)
	}

	dst, err := url.Parse(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var origin string
	if c.url != "" {
		origin = c.url
	} else {
		originURL := *dst
		if dst.Scheme == "wss" {
			originURL.Scheme = "https"
		} else {
			originURL.Scheme = "http"
		}
		origin = originURL.String()
	}

	var historyFile string
	u, err := user.Current()
	if err == nil {
		historyFile = filepath.Join(u.HomeDir, ".websocket_cli_history")
	}
	token := "Bearer " + c.token
	err = ws.Connect(dst.String(), origin, token, &readline.Config{
		Prompt:      "> ",
		HistoryFile: historyFile,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if err != io.EOF && err != readline.ErrInterrupt {
			os.Exit(1)
		}
	}

	return nil
}