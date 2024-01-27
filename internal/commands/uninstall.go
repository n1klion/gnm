package commands

import (
	"errors"
	"fmt"
	"gnm/internal/node"
	"gnm/internal/version"
	"gnm/utils/config"

	"github.com/urfave/cli/v2"
)

func Uninstall(c *cli.Context) error {
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

	n := node.NewNode(cfg, vsn)
	if err := n.Uninstall(); err != nil {
		return fmt.Errorf("failed to remove Node.js version %s: %v", vsn.String(), err)
	}

	fmt.Printf("Node.js version %s uninstalled successfully.\n", vsn.String())

	return nil
}
