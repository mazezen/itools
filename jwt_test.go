package itools

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var ex = time.Minute * 10
var sec = "70fc4956fb1d4f0f963d9cc2a3bf86e934234223425"

func TestJwtClaims_GenerateToken(t *testing.T) {

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
}

func TestJwtClaims_ParseToken(t *testing.T) {

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmUiOjAsIlNlY3JldCI6IiIsImxvZ2luX2luZm8iOnsiaWQiOjEsIm5hbWUiOiJhZG1pbiJ9LCJleHAiOjE3MjY3MTYxODZ9.16OrV7p3vDRuf2mc1_iJnHWc3RfqevJNoD5JSTDNQtw"
	js := NewJwt(ex, sec)

	info, err := js.ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	var loginInfo struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	bytes, _ := json.Marshal(info.LoginInfo)
	_ = json.Unmarshal(bytes, &loginInfo)

	fmt.Println(loginInfo.Id)
	fmt.Println(loginInfo.Name)
}
