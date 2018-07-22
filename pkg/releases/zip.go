package releases

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func uncompressZip(archive, file, destination string) error {
	zipReader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	var zipFile *zip.File
	for _, f := range zipReader.File {
		if f.FileHeader.Name == file {
			zipFile = f
			break
		}
	}
	if zipFile == nil {
		return fmt.Errorf("Cannot find file %s in archive", file)
	}

	dst, err := os.OpenFile(
		destination,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0755,
	)
	if err != nil {
		return err
	}
	defer dst.Close()

	src, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = io.Copy(dst, src)
	return err
}
