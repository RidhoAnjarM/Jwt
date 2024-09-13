package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func Upload(location string, file multipart.File, header *multipart.FileHeader) string {
	exts := strings.Split(header.Filename, ".")
	filename := UUID() + "." + exts[len(exts)-1]
	dir := "static/uploads/" + location + "/"

	_ = os.MkdirAll(dir, os.ModePerm)
	dst, _ := os.Create(dir + filename)

	_, _ = io.Copy(dst, file)
	_ = dst.Close()

	return dir + filename
}
