package learn_jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var hmacSecret = []byte("fms5cmS")

// Header 部分：alg=HS256、typ=JWT
// Payload(也叫 Claims)部分：
//   Registered 部分这里未设置
//   Public 部分需要额外注册？
//   Private 部分可以随意定义，这里的两个字段都是 Private 的
// Signature 部分，防止数据被篡改
//   注意，该库的 SignedString 方法接收一个 interface{} 类型，但必须是 []byte，否则运行时报错
func genToken(username, curTime string) (string, error) {
	// Claims 用于存储数据
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"login-time": curTime,
	})
	return token.SignedString(hmacSecret)
}

func parseAndValidateToken(jwtToken string) (map[string]string, error) {
	// 解析 token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 对 alg 即签名算法校验
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v ", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	// 校验有效性，并获取 Claims 中的值
	// Valid 验证基于时间的声明，如 exp, iat, nbf，注意如果在令牌中没有任何声明，仍然会被认为是有效的。并且对于时区偏差没有计算方法
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		loginTime := claims["login-time"].(string)
		return map[string]string{
			"username":   username,
			"login-time": loginTime,
		}, nil
	}
	return nil, err
}
