package itools

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"Md5s":    {testMd5S},
		"Sha1":    {testSha1},
		"FileMd5": {testFileMd5},
		"Sha256":  {testSha256},
	}
	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

func testMd5S(t *testing.T) {
	fmt.Println(Md5S("123456"))
}

func testSha1(t *testing.T) {
	fmt.Println(Sha1("123456"))
}

func testFileMd5(t *testing.T) {
	fmt.Println(FileMd5("./aes.go"))
}

func testSha256(t *testing.T) {
	fmt.Println(Sha256("123456"))
}
