package version

import (
	"fmt"
	"gnm/internal/constants"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Version struct {
	major int
	minor int
	patch int
}

func NewVersion() *Version {
	return &Version{}
}

func (v *Version) TryParse(inputVersion string) error {
	vs := strings.Split(inputVersion, ".")

	switch len(vs) {
	case 3:
		patch, err := strconv.Atoi(vs[2])
		if err != nil {
			return fmt.Errorf("failed to parse patch version: %v", err)
		}
		v.patch = patch
		fallthrough
	case 2:
		minor, err := strconv.Atoi(vs[1])
		if err != nil {
			return fmt.Errorf("failed to parse minor version: %v", err)
		}
		v.minor = minor
		fallthrough
	case 1:
		major, err := strconv.Atoi(vs[0])
		if err != nil {
			return fmt.Errorf("failed to parse major version: %v", err)
		}
		v.major = major
	default:
		return fmt.Errorf("invalid version format: %s", inputVersion)
	}

	return nil
}

func (v *Version) GetDownloadURL() string {
	return fmt.Sprintf("%s/v%s/node-v%s-%s-%s.tar.gz", constants.DownloadURL, v.String(), v.String(), getOS(), getArch())
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func (v *Version) ParseEqualOrMajor(version string) error {
	installed, err := GetInstalled()
	if err != nil {
		return fmt.Errorf("failed to get installed Node.js versions: %v", err)
	}

	for _, i := range installed {
		if i == version {
			if err := v.TryParse(i); err != nil {
				return fmt.Errorf("failed to parse version: %v", err)
			}
			break
		}
	}

	for _, i := range installed {
		if strings.HasPrefix(i, version) {
			if err := v.TryParse(i); err != nil {
				return fmt.Errorf("failed to parse version: %v", err)
			}
			break
		}
	}

	return nil
}

func (v *Version) IsCurrentInUse(versionFromConfig string) (bool, error) {
	if versionFromConfig == "" {
		return false, nil
	}

	if versionFromConfig == v.String() {
		return true, nil
	}

	return false, nil
}

func GetInstalled() ([]string, error) {
	versions := []string{}
	files, err := os.ReadDir(filepath.Join(os.Getenv("HOME"), constants.GnmDirName, constants.VersionDirName))
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			versions = append(versions, file.Name())
		}
	}

	return versions, nil
}

func getOS() string {
	switch runtime.GOOS {
	case "darwin":
		return "darwin"
	case "linux":
		return "linux"
	case "windows":
		return "win"
	default:
		return "linux"
	}
}

func getArch() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x64"
	case "386":
		return "x86"
	case "arm":
		return "armv6l"
	default:
		return "x64"
	}
}
