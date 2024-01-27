package commands

import (
	"fmt"
	"gnm/internal/constants"
	"gnm/internal/node"
	"gnm/internal/version"
	"gnm/utils/config"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"
)

func SessionStart(c *cli.Context) error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	versionFromArgs := cfg.GetDefault()
	vsn := version.NewVersion()
	if err := vsn.ParseEqualOrMajor(versionFromArgs); err != nil {
		return fmt.Errorf("not found Node.js version %s", versionFromArgs)
	}

	n := node.NewNode(cfg, vsn)
	symlinkPath, err := n.SetCurrent()
	if err != nil {
		return fmt.Errorf("failed to switch to Node.js version %s: %v", vsn.String(), err)
	}

	fmt.Print(symlinkPath)

	return nil
}

func SessionEnd(c *cli.Context) error {
	pathToSymlink := filepath.Join(os.Getenv("HOME"), constants.GnmDirName, "tty", strconv.Itoa(os.Getppid()))
	if err := node.NewNode(nil, nil).UnlinkCurrent(pathToSymlink); err != nil {
		return fmt.Errorf("failed to unlink current version: %v", err)
	}

	return nil
}
