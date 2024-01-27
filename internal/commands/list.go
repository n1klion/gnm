package commands

import (
	"fmt"
	"gnm/internal/version"
	"gnm/utils/config"

	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	versions, err := version.GetInstalled()
	if err != nil {
		return fmt.Errorf("failed to get installed Node.js versions: %v", err)
	}

	if len(versions) == 0 {
		fmt.Println("No Node.js versions installed.")
		return nil
	}

	fmt.Println("Installed Node.js versions:")
	for _, file := range versions {
		if file == cfg.GetCurrent() {
			fmt.Printf("\u001b[32m%s (current)\u001b[0m\n", file)
		} else {
			fmt.Println(file)
		}
	}

	return nil
}
