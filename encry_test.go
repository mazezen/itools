package itools

import (
	"fmt"
	"testing"
)

func TestMd5S(t *testing.T) {
	fmt.Println(Md5S("123456"))
}

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123456"))
}

func TestFileMd5(t *testing.T) {
	fmt.Println(FileMd5("./aes.go"))
}
