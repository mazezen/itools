package itools

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"testJwtClaims_GenerateToken": {JwtClaimsGenerateToken},
	}
	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

var ex = time.Minute * 10
var sec = "70fc4956fb1d4f0f963d9cc2a3bf86e934234223425"

func JwtClaimsGenerateToken(t *testing.T) {

	var loginInfo struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
	loginInfo.Id = 1
	loginInfo.Name = "admin"

	js := NewJwt(ex, sec)
	token, err := js.GenerateToken(loginInfo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)

	info, err := js.ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}

	bytes, _ := json.Marshal(info.LoginInfo)
	_ = json.Unmarshal(bytes, &loginInfo)

	fmt.Println(loginInfo.Id)
	fmt.Println(loginInfo.Name)
}
