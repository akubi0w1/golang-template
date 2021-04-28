package jwt

import (
	"time"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
)

// TODO: key方式にしてください
const secretKey string = "sample"

type TokenHandler interface {
	Generate(accountID string) (string, error)
	Parse(tokenString string) (Claims, error)
}

type JWTHandler struct{}

func NewJWTHandler() *JWTHandler {
	return &JWTHandler{}
}

func (h *JWTHandler) Generate(accountID string) (string, error) {
	claims := Claims{
		AccountID: accountID,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", code.Errorf(code.JWT, "failed to generate token: %v", err)
	}

	return tokenString, nil
}

func (h *JWTHandler) Parse(tokenString string) (Claims, error) {
	token, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return Claims{}, code.Errorf(code.Unauthorized, "failed to parse token: %v", err)
	}

	claims := toClaims(token.Claims.(jwt.MapClaims))
	return claims, nil
}

type Claims struct {
	AccountID      string               `json:"accountId"`
	StandardClaims jwtgo.StandardClaims `json:"standardClaims"`
}

func (c Claims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return code.Errorf(code.BadRequest, "failed to validate token claims: %v", err)
	}
	return nil
}

func (c *Claims) GetAccountID() string {
	return c.AccountID
}

func toClaims(c jwtgo.MapClaims) Claims {
	standardClaims := c["standardClaims"].(map[string]interface{})
	return Claims{
		AccountID: c["accountId"].(string),
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: int64(standardClaims["exp"].(float64)),
		},
	}
}
