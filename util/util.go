package util

import (
	"crypto/md5" 
	"encoding/hex" 
	"time"
	"strconv"
	"path/filepath"
	"strings"
	"log"
	"os"
)
// 加密密码
func EncryptionPassword(old string) string {
	h := md5.New()
	h.Write([]byte(old))
	new := h.Sum(nil)
	return hex.EncodeToString(new)
}
// 根据时间获取Id
func GetRowId() string {
	old := time.Now().Nanosecond()
	h := md5.New()
	h.Write([]byte(strconv.Itoa(old)))
	new := h.Sum(nil)
	return hex.EncodeToString(new)
}
// 获取当前程序运行的绝对路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(dir)
	return strings.Replace(dir, "\\", "/", -1)
}