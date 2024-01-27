package gnm

import (
	"gnm/internal/commands"

	"github.com/urfave/cli/v2"
)

func Initialize() *cli.App {
	app := &cli.App{
		Name:  "gnm",
		Usage: "A simple package manager for Node.js",
		Commands: []*cli.Command{
			{
				Name:    "install",
				Usage:   "Install a package",
				Aliases: []string{"i"},
				Action:  commands.Install,
			},
			{
				Name:    "list",
				Usage:   "List installed packages",
				Aliases: []string{"ls"},
				Action:  commands.List,
			},
			{
				Name:    "uninstall",
				Usage:   "Uninstall a package",
				Aliases: []string{"un", "rm"},
				Action:  commands.Uninstall,
			},
			{
				Name:   "use",
				Usage:  "Use a package",
				Action: commands.Use,
			},
			{
				Name:  "session",
				Usage: "Manage sessions",
				Subcommands: []*cli.Command{
					{
						Name:   "new",
						Usage:  "Start a new session",
						Action: commands.SessionStart,
					},
					{
						Name:   "close",
						Usage:  "Close the current session",
						Action: commands.SessionEnd,
					},
				},
			},
		},
	}

	return app
}
