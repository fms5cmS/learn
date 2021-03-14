package learn_jwt

import (
	"encoding/base64"
	"testing"
	"time"
)

func TestJwtToken(t *testing.T) {
	token, err := genToken("zzk", time.Now().String())
	if err != nil {
		t.Error("generate token error: ", err)
	}
	t.Log(token)
	info, err := parseAndValidateToken(token)
	if err != nil {
		t.Error("parse token error: ", err)
	}
	t.Log(info)
}

// 将得到的 token 中 payload 复制后进行解析
func TestJWT(t *testing.T) {
	payload,_ := base64.StdEncoding.DecodeString("eyJhcHBfa2V5IjoiZm1zNWNtUyIsImFwcF9zZWNyZXQiOiJtczVjbWthaSIsImV4cCI6MTYwMTc4Nzk2NiwiaXNzIjoiYmxvZy1l\ncnZpY2UifQ")
	t.Log(string(payload))
}