package auth

import (
	"github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xssnick/tonutils-go/address"
)

const secret = "secret-jwt-ton-in-mon"

type JwtClaims struct {
	AccountId string           `json:"account_id"`
	Address   *address.Address `json:"address"`
	Ts        int64            `json:"ts"`
	Domain    string           `json:"domain"`
	Exp       int64            `json:"exp"`
}

func CreateJwtToken(claims jwt.MapClaims) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(secret))
}

func JwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	})
}

func ConvertMapClaims(claims jwt.MapClaims) *JwtClaims {
	return &JwtClaims{
		AccountId: claims["account_id"].(string),
		Address:   address.MustParseRawAddr(claims["account_id"].(string)),
		Ts:        int64(claims["ts"].(float64)),
		Domain:    claims["domain"].(string),
		Exp:       int64(claims["exp"].(float64)),
	}
}
