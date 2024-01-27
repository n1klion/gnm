package commands

import (
	"fmt"
	"gnm/utils/config"
	"os"

	"github.com/urfave/cli/v2"
)

func Env(c *cli.Context) error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}
	currentVersion := cfg.GetCurrent()
	if currentVersion == "" {
		fmt.Println("No default Node.js version set")
		return nil
	}

	return updateCurrentEnv(currentVersion)
}

func updateCurrentEnv(version string) error {
	return os.Setenv("PATH", fmt.Sprintf("%s:~/.gnm/versions/%s/bin", os.Getenv("PATH"), version))
}
