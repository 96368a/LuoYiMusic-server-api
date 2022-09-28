package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func GetFileMd5(file *multipart.FileHeader) string {
	fileContent, _ := file.Open()
	byteContainer, _ := io.ReadAll(fileContent)
	hash := md5.Sum(byteContainer)
	md5str := fmt.Sprintf("%x", hash) //将[]byte转成16进制
	return md5str
}
