package itools

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"AesEncryptDecrypt": {testAesEncryptDecrypt},
	}
	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

func testAesEncryptDecrypt(t *testing.T) {
	instance, err := NewAesEncryptInstance("bGcGfWb3Kg2s4gcG", "aebksHkG4jAEk2Ag")
	if err != nil {
		t.Fatal(err)
	}
	encrypt, err := instance.AesBase64Encrypt("1234566dfsdasd")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encrypt)

	fmt.Println("=======================================")

	decrypt, err := instance.AesBase64Decrypt(encrypt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypt)
}
