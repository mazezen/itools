package itools

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"Md5encode":     {testMd5encoder},
		"Md5StrToUpper": {testMd5StrToUpper},
		"Md5SaltCode":   {testMd5SaltCode},
		"Sha1":          {testSha1},
		"FileMd5":       {testFileMd5},
		"Sha256":        {testSha256},
	}
	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

func testMd5encoder(t *testing.T) {
	fmt.Println(Md5encoder("123456"))
}

func testMd5StrToUpper(t *testing.T) {
	t.Logf("Md5StrToUpper: %s", Md5StrToUpper("123456"))
}

func testMd5SaltCode(t *testing.T) {
	slat := rand.Int31()
	s, _ := ToString(slat)
	encoder := Md5SaltCode("123456", s)
	fmt.Println(encoder)
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
