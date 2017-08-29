package main

import (
	sh "github.com/ipfs/go-ipfs-api"
	"time"
	"BotCommands"
)

func main() {
	shell := sh.NewShell("https://ipfs.io")

	ticker := time.NewTicker(time.Second * 10)
	go BotCommands.FetchCommand(ticker, shell)

	select {}
}
