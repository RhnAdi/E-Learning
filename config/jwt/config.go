package jwt

import (
	"fmt"
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

func NewJWTConfig(rootPath string) (cfg JWTConfig, err error) {
	cfg.AccessTokenPrivateKey, err = os.ReadFile(fmt.Sprintf("%saccess_private_key.pub", rootPath))
	if err != nil {
		return
	}
	cfg.AccessTokenPublicKey, err = os.ReadFile(fmt.Sprintf("%saccess_public_key.pub", rootPath))
	if err != nil {
		return
	}
	cfg.RefreshTokenPrivateKey, err = os.ReadFile(fmt.Sprintf("%srefresh_private_key.pub", rootPath))
	if err != nil {
		return
	}
	cfg.RefreshTokenPublicKey, err = os.ReadFile(fmt.Sprintf("%srefresh_public_key.pub", rootPath))
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
