package node

import (
	"fmt"
	"gnm/internal/constants"
	"gnm/internal/downloader"
	"gnm/internal/extractor"
	"gnm/internal/version"
	"gnm/utils/config"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

type Node struct {
	version   *version.Version
	appConfig *config.Config
}

func NewNode(cfg *config.Config, version *version.Version) *Node {
	return &Node{
		version:   version,
		appConfig: cfg,
	}
}

func (n *Node) SetCurrent() (string, error) {
	if err := n.appConfig.UpdateCurrent(n.version.String()); err != nil {
		return "", fmt.Errorf("failed to update current Node.js version: %v", err)
	}

	symlinksLocation := filepath.Join(os.Getenv("HOME"), constants.GnmDirName, "tty")
	if _, err := os.Stat(symlinksLocation); os.IsNotExist(err) {
		if err := os.MkdirAll(symlinksLocation, 0755); err != nil {
			return "", fmt.Errorf("failed to create symlinks directory: %v", err)
		}
	}

	symlinkPath := filepath.Join(symlinksLocation, strconv.Itoa(os.Getppid()))
	if err := n.UnlinkCurrent(symlinkPath); err != nil {
		return "", fmt.Errorf("failed to unlink current version: %v", err)
	}

	curBinPath := filepath.Join(os.Getenv("HOME"), constants.GnmDirName, constants.VersionDirName, n.version.String(), "bin")
	if err := os.Symlink(curBinPath, symlinkPath); err != nil {
		return "", fmt.Errorf("failed to create symlink: %v", err)
	}
	return symlinkPath, nil
}

func (n *Node) Install() error {
	installed, err := n.IsInstalled()
	if err != nil {
		return err
	}
	if installed {
		fmt.Printf("Node.js version %s is already installed\n", n.version.String())
		return nil
	}

	archivePath := fmt.Sprintf("%s.%s", filepath.Join(os.Getenv("HOME"), constants.GnmDirName, constants.ArchivesDirName, n.version.String()), "tar.gz")
	dl := downloader.NewDownloader(n.version.GetDownloadURL(), archivePath)
	if err := dl.Download(n.installProgressHandler); err != nil {
		return err
	}

	installationPath := filepath.Join(os.Getenv("HOME"), constants.GnmDirName, constants.VersionDirName, n.version.String())
	ext := extractor.NewExtractor(archivePath, installationPath)
	if err = ext.Extract(); err != nil {
		return err
	}

	if _, err := n.SetCurrent(); err != nil {
		return err
	}

	return nil
}

func (n *Node) Uninstall() error {
	pathToVersion := filepath.Join(os.Getenv("HOME"), constants.GnmDirName, constants.VersionDirName, n.version.String())
	if _, err := os.Stat(pathToVersion); os.IsNotExist(err) {
		return nil
	}

	if err := os.RemoveAll(pathToVersion); err != nil {
		return err
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}
	if err = cfg.UpdateCurrent(""); err != nil {
		return fmt.Errorf("failed to update current Node.js version: %v", err)
	}

	return nil
}

func (n *Node) IsInstalled() (bool, error) {
	pathToVersion := filepath.Join(os.Getenv("HOME"), constants.GnmDirName, constants.VersionDirName, n.version.String())
	if _, err := os.Stat(pathToVersion); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func (n *Node) UnlinkCurrent(symlinkPath string) error {
	if _, err := os.Lstat(symlinkPath); err == nil {
		if err := os.Remove(symlinkPath); err != nil {
			return fmt.Errorf("failed to remove old symlink: %v", err)
		}
	}

	return nil
}

func (n *Node) installProgressHandler(downloaded, total int64) {
	elementsCount := 20
	progressPercent := int(downloaded * 100 / total)

	drawProgressTo := int(math.Round(float64(progressPercent*elementsCount) / 100))
	var progressBar string
	for i := 0; i < elementsCount; i++ {
		if i < drawProgressTo {
			progressBar += "#"
		} else {
			progressBar += "_"
		}
	}

	fmt.Printf("\rDownloading Node.js version %s [%s] %d%%", n.version.String(), progressBar, progressPercent)
}
