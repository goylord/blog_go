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
func EncryptionPassword(old string) string {
	h := md5.New()
	h.Write([]byte(old))
	new := h.Sum(nil)
	return hex.EncodeToString(new)
}
func GetRowId() string {
	old := time.Now().Nanosecond()
	h := md5.New()
	h.Write([]byte(strconv.Itoa(old)))
	new := h.Sum(nil)
	return hex.EncodeToString(new)
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(dir)
	return strings.Replace(dir, "\\", "/", -1)
}