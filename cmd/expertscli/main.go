package main

import (
	"fmt"

	"github.com/leodev0/expertscli/commands/message"
	"github.com/leodev0/expertscli/commands/start"
	"github.com/leodev0/expertscli/internal"
)

var commandList = []internal.Command{
	new(start.Start),
	new(message.Message),
}

func main() {
	err := internal.CommandInit("expertscli").Start(commandList)

	if err != nil {
		fmt.Println(err.Error())
	}
}