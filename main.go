package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC

	file, err := os.OpenFile(os.Args[-1], flags, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for _, filename := range os.Args[1:] {
		if err := AddFileToZip(zipWriter, filename); err != nil {
			panic(err)
		}
	}

}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	splitted_filename := strings.Split(filename, "/")
	filename_new := splitted_filename[len(splitted_filename)-1]

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename_new

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
