package BotCommands

import (
	"regexp"
	"strconv"
	"strings"
	"errors"
)

type BotCommand struct {
	Index int
	MainCommand string
	Arguments []string
}

func ParseBotCommand(rawCommand string) (*BotCommand, error) {
	re := regexp.MustCompile(`(\d+)\s(\w+)(?:\s(.+))?`)
	sm := re.FindStringSubmatch(rawCommand)

	if sm == nil {
		return nil, errors.New("Invalid raw command")
	}

	index, _ := strconv.Atoi(sm[1])
	mainCommand := sm[2]
	arguments := strings.Split(sm[3], " ")

	return &BotCommand{index, mainCommand, arguments}, nil
}
