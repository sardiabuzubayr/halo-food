package security

import (
	"halo_food/config"

	"github.com/golang-jwt/jwt"
)

type (
	JwtToken struct {
		Email     string `json:"email"`
		Alias     string `json:"alias"`
		Nama      string `json:"nama"`
		IdLevel   int    `json:"id_level"`
		NamaLevel string `json:"nama_level"`
		jwt.StandardClaims
	}

	JwtRefreshToken struct {
		Email string `json:"email"`
		Alias string `json:"alias"`
		jwt.StandardClaims
	}
)

func EncodeToken(data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenResult, err := token.SignedString([]byte(config.GetEnv("TOKEN_SECRET")))
	return tokenResult, err
}
