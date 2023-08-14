package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrCopyToYouself         = errors.New("copy to yourself")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == toPath {
		return ErrCopyToYouself
	}

	if limit == 0 {
		byteBuffer, err := os.ReadFile(fromPath)
		if err != nil {
			return err
		}

		err = os.WriteFile(toPath, byteBuffer, 0o600)
		if err != nil {
			return err
		}

		return nil
	}

	inputFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	targetFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	fileInfo, err := inputFile.Stat()
	if fileInfo.IsDir() {
		return ErrUnsupportedFile
	}
	if err != nil {
		return err
	}
	size := fileInfo.Size()

	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	if offset > 0 {
		_, err = inputFile.Seek(offset, 0)
		if err != nil {
			return err
		}
	}

	bar := pb.Simple.Start64(limit)
	defer bar.Finish()
	barReader := bar.NewProxyReader(inputFile)

	if _, err = io.CopyN(targetFile, barReader, limit); err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}
