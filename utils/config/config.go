package config

import (
	"encoding/json"
	"errors"
	"gnm/internal/constants"
	"os"
	"path/filepath"
)

type configJSON struct {
	Current string `json:"current"`
	Default string `json:"default"`
}

type Config struct {
	currentNode string
	defaultNode string
}

func ReadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	configFileRaw, err := os.ReadFile(configPath)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			return &Config{}, nil
		default:
			return nil, err
		}
	}

	cfgJSON := &configJSON{}
	if err := json.Unmarshal(configFileRaw, cfgJSON); err != nil {
		return nil, err
	}

	cfg := &Config{
		currentNode: cfgJSON.Current,
		defaultNode: cfgJSON.Default,
	}

	return cfg, nil
}

func (c *Config) UpdateCurrent(version string) error {
	c.currentNode = version
	return c.SaveConfig()
}

func (c *Config) UpdateDefault(version string) error {
	c.defaultNode = version
	return c.SaveConfig()
}

func (c *Config) GetCurrent() string {
	return c.currentNode
}

func (c *Config) GetDefault() string {
	return c.defaultNode
}

func (c *Config) SaveConfig() error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	cfgJSON := &configJSON{
		Current: c.currentNode,
		Default: c.defaultNode,
	}

	configFileRaw, err := json.Marshal(cfgJSON)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configPath, configFileRaw, 0644); err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return "", errors.New("failed to get home directory")
	}

	return filepath.Join(homeDir, constants.GnmDirName, "config.json"), nil
}
