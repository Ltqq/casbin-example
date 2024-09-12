package midware

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/maxwellhertz/gin-casbin"
)

var Auth *gcasbin.CasbinMiddleware

func InitCasbinAuth() error {
	// Use Casbin authentication middleware.
	auth, err := gcasbin.NewCasbinMiddleware("./conf/model.conf", "./conf/policy.csv", subjectFromJWT)
	if err != nil {
		return err
	}
	Auth = auth
	return nil
}

// subjectFromJWT parses a JWT and extract subject from sub claim.
func subjectFromJWT(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	token := authHeader
	if token == "" {
		// JWT not found.
		return ""
	}

	var payload jwt.Payload
	_, err := jwt.Verify([]byte(token), jwtKey, &payload)
	if err != nil {
		return ""
	}
	return payload.Subject
}
