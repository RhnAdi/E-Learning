package jwt

import (
	"os"
	"strconv"
	"time"
)

type JWTClaims struct {
	Id   string
	Role string
}

type JWTConfig struct {
	AccessTokenPrivateKey  []byte
	AccessTokenPublicKey   []byte
	RefreshTokenPrivateKey []byte
	RefreshTokenPublicKey  []byte
	AccessTokenMaxAge      int
	RefreshTokenMaxAge     int
	AccessTokenExpiresIn   time.Duration
	RefreshTokenExpiresIn  time.Duration
}

func NewJWTConfig() (cfg JWTConfig, err error) {
	cfg.AccessTokenPrivateKey, err = os.ReadFile("../../../access_private_key.pub")
	if err != nil {
		return
	}
	cfg.AccessTokenPublicKey, err = os.ReadFile("../../../access_public_key.pub")
	if err != nil {
		return
	}
	cfg.RefreshTokenPrivateKey, err = os.ReadFile("../../../refresh_private_key.pub")
	if err != nil {
		return
	}
	cfg.RefreshTokenPublicKey, err = os.ReadFile("../../../refresh_public_key.pub")
	if err != nil {
		return
	}

	cfg.AccessTokenMaxAge, err = strconv.Atoi(os.Getenv("ACCESS_TOKEN_MAXAGE"))
	if err != nil {
		return
	}
	cfg.RefreshTokenMaxAge, err = strconv.Atoi(os.Getenv("REFRESH_TOKEN_MAXAGE"))
	if err != nil {
		return
	}
	cfg.AccessTokenExpiresIn = time.Duration(cfg.AccessTokenMaxAge) * time.Minute
	cfg.RefreshTokenExpiresIn = time.Duration(cfg.RefreshTokenMaxAge) * time.Minute

	return
}
