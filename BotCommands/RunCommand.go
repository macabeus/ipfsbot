package BotCommands

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func (botCommand *BotCommand) RunCommand() {
	switch botCommand.MainCommand {
	case "get":
		fmt.Println("[run command] get with ", botCommand.Arguments)
		commandGetAndRun(botCommand.Arguments[0])

	default:
		fmt.Println("I didn't know this commando, sorry")
	}
}

func commandGetAndRun(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("[get command] error: ", err)
		return
	}

	defer response.Body.Close()

	urlSplit := strings.Split(url, "/")
	fileName := urlSplit[len(urlSplit) - 1]

	out, err := os.Create(os.TempDir() + fileName)
	if err != nil {
		fmt.Println("[get command] error: ", err)
		return
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		fmt.Println("[get command] error: ", err)
		return
	}

	fmt.Println("[get command] success")
}
