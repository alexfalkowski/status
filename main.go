package main

import (
	"os"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/status/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *scmd.Command {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)
	command.AddVersion(cmd.Version)

	return command
}
