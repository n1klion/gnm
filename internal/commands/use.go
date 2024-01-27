package commands

import (
	"errors"
	"fmt"
	"gnm/internal/node"
	"gnm/internal/version"
	"gnm/utils/config"

	"github.com/urfave/cli/v2"
)

func Use(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("no Node.js version specified")
	}
	versionFromArgs := c.Args().First()

	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}
	vsn := version.NewVersion()
	if err := vsn.ParseEqualOrMajor(versionFromArgs); err != nil {
		return fmt.Errorf("not found Node.js version %s", versionFromArgs)
	}

	n := node.NewNode(cfg, vsn)
	if _, err := n.SetCurrent(); err != nil {
		return fmt.Errorf("failed to switch to Node.js version %s: %v", vsn.String(), err)
	}

	fmt.Printf("Switched to Node.js version %s.\n", vsn.String())

	return nil
}
