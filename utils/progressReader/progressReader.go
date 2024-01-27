package progressReader

import "io"

type ProgressReader struct {
	io.Reader
	total      int64
	downloaded int64
	onProgress func(downloaded, total int64)
}

func New(r io.Reader, total int64, onProgress func(downloaded, total int64)) io.Reader {
	return &ProgressReader{
		Reader:     r,
		total:      total,
		onProgress: onProgress,
	}
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.downloaded += int64(n)
	pr.onProgress(pr.downloaded, pr.total)
	return n, err
}
