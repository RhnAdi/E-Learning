package security_test

import (
	"testing"

	"github.com/RhnAdi/elearning-microservice/security"
	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	pass, err := security.EcryptPassword("123456789")

	assert.NoError(t, err)
	assert.NotEmpty(t, pass)
	assert.Len(t, pass, 60)
}

func TestVerifyPassword(t *testing.T) {
	pass, err := security.EcryptPassword("123456789")

	assert.NoError(t, err)
	assert.NotEmpty(t, pass)
	assert.Len(t, pass, 60)

	assert.NoError(t, security.VerifyPassword(pass, "123456789"))
	assert.Error(t, security.VerifyPassword(pass, "wrong password"))
	assert.Error(t, security.VerifyPassword("123456789", pass))
	assert.Error(t, security.VerifyPassword(pass, pass))
}
