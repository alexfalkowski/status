package cmd

import (
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// RegisterServer adds the server subcommand backed by Module and its shared configuration flags.
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start status server", Module)

	cmd.AddConfig(strings.Empty)
}
