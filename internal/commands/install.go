package commands

import (
	"errors"
	"fmt"

	"gnm/internal/node"
	"gnm/internal/version"
	"gnm/utils/config"

	"github.com/urfave/cli/v2"
)

func Install(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("no package specified")
	}

	versionFromArgs := c.Args().First()
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}
	vsn := version.NewVersion()
	if err = vsn.TryParse(versionFromArgs); err != nil {
		return err
	}

	fmt.Printf("Installing Node.js version %s...\n", vsn.String())

	n := node.NewNode(cfg, vsn)
	if err = n.Install(); err != nil {
		return fmt.Errorf("failed to install Node.js version %s: %v", vsn.String(), err)
	}

	fmt.Printf("\nNode.js version %s installed successfully.", vsn.String())

	return nil
}
