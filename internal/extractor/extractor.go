package extractor

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Extractor struct {
	archivePath string
	outputPath  string
}

func NewExtractor(archivePath, outputPath string) *Extractor {
	return &Extractor{
		archivePath: archivePath,
		outputPath:  outputPath,
	}
}

func (e *Extractor) Extract() error {
	archStream, err := os.Open(e.archivePath)
	if err != nil {
		return err
	}

	gzipStream, err := gzip.NewReader(archStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzipStream)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(e.outputPath, removeRoot(header.Name))
		needExecuteRules := strings.HasSuffix(target, "node") || strings.HasSuffix(target, "npm") || strings.HasSuffix(target, "npx")

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, target); err != nil {
				return err
			}
			if needExecuteRules {
				if err := os.Chmod(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			outFile, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()

			if needExecuteRules {
				if err := os.Chmod(target, 0755); err != nil {
					return err
				}
			}
		}
	}

	os.Remove(e.archivePath)

	return nil
}

func removeRoot(path string) string {
	return strings.Join(strings.Split(path, "/")[1:], "/")
}
