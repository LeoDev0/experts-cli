package main

import (
	"fmt"

	"github.com/leodev0/expertscli/commands"
	"github.com/leodev0/expertscli/internal"
)

var commandList = []internal.Command{
	new(commands.Start),
}

func main() {
	err := internal.CommandInit("expertscli").Start(commandList)

	if err != nil {
		fmt.Println(err.Error())
	}
}