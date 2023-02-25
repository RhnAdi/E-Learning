package security_test

import (
	"testing"

	"github.com/RhnAdi/elearning-microservice/config/jwt"
	"github.com/RhnAdi/elearning-microservice/security"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	err := godotenv.Load("../cmd/grpc/auth/.env")
	assert.NoError(t, err)
	config, err := jwt.NewJWTConfig("../")
	assert.NoError(t, err)
	token, err := security.CreateToken(config.AccessTokenExpiresIn, jwt.JWTClaims{
		Id:   "123",
		Role: "teacher",
	}, config.AccessTokenPrivateKey)
	assert.NoError(t, err)
	res, err := security.ValidateToken(token, config.AccessTokenPublicKey)
	assert.NoError(t, err)
	sub := res.(map[string]interface{})
	claim := jwt.JWTClaims{
		Id:   sub["Id"].(string),
		Role: sub["Role"].(string),
	}
	t.Logf("claim.Id : %s claim.Role : %s", claim.Id, claim.Role)
}
