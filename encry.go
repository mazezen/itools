package itools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func Md5encoder(code string) string {
	w := md5.New()
	_, _ = io.WriteString(w, code)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	return md5Str
}

func Md5StrToUpper(code string) string {
	return strings.ToUpper(Md5encoder(code))
}

// Md5SaltCode 加密加盐
func Md5SaltCode(code, slat string) string {
	return CompactStr(Md5encoder(code), slat)
}

func Sha1(code string) string {
	o := sha1.New()
	o.Write([]byte(code))
	return hex.EncodeToString(o.Sum(nil))
}

func FileMd5(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Sha256(code string) string {
	hash := sha256.New()
	hash.Write([]byte(code))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}
