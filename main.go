package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/status/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:CONFIG_FILE")
	c.AddServer("server", "Start status server", cmd.ServerOptions...)

	return c
}
