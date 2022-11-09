package file

import "strings"

// SeparateFileSuffix 分离文件后缀名
func SeparateFileSuffix(fileName string) string {
	return fileName[strings.LastIndex(fileName, "."):]
}
