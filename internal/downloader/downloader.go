package downloader

import (
	"fmt"
	"gnm/utils/progressReader"
	"io"
	"net/http"
	"os"
)

type Downloader struct {
	downloadURL string
	savePath    string
}

func NewDownloader(downloadURL string, savePath string) *Downloader {
	return &Downloader{
		downloadURL,
		savePath,
	}
}

func (d *Downloader) Download(onDownload func(downloaded, total int64)) error {
	resp, err := http.Get(d.downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}
	defer resp.Body.Close()

	file, err := os.Create(d.savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	pr := progressReader.New(resp.Body, resp.ContentLength, onDownload)

	_, err = io.Copy(file, pr)
	if err != nil {
		return fmt.Errorf("failed to copy: %v", err)
	}

	return nil
}
