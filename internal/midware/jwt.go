package midware

import (
	"github.com/gbrlsnchs/jwt/v3"
	"time"
)

var jwtKey = jwt.NewHS256([]byte("jwtkey"))

func GetToken(uname string) (string, error) {
	payload := jwt.Payload{
		Subject:        uname,
		ExpirationTime: jwt.NumericDate(time.Now().Add(time.Hour)),
	}
	token, err := jwt.Sign(payload, jwtKey)
	if err != nil {
		return "", err
	}
	return string(token), nil
}
