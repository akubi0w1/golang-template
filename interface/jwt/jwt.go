package jwt

import (
	"time"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
)

// TODO: key方式にしてください
const secretKey string = "sample"

type JWTImpl struct{}

func NewJWTImpl() *JWTImpl {
	return &JWTImpl{}
}

func (h *JWTImpl) Generate(claims entity.Claims) (entity.Token, error) {
	_claims := jwtClaims{
		AccountID: claims.AccountID,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, _claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", code.Errorf(code.JWT, "failed to generate token: %v", err)
	}

	return entity.Token(tokenString), nil
}

func (h *JWTImpl) Parse(token entity.Token) (entity.Claims, error) {
	_token, err := jwtgo.Parse(token.String(), func(token *jwtgo.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return entity.Claims{}, code.Errorf(code.Unauthorized, "failed to parse token: %v", err)
	}

	claims := toEntityClaims(_token.Claims.(jwt.MapClaims))
	return claims, nil
}

type jwtClaims struct {
	AccountID      entity.AccountID     `json:"accountId"`
	StandardClaims jwtgo.StandardClaims `json:"standardClaims"`
}

func (c jwtClaims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return code.Errorf(code.BadRequest, "failed to validate token claims: %v", err)
	}
	return nil
}

func toEntityClaims(c jwtgo.MapClaims) entity.Claims {
	// standardClaims := c["standardClaims"].(map[string]interface{})
	return entity.Claims{
		AccountID: entity.AccountID(c["accountId"].(string)),
		// StandardClaims: jwtgo.StandardClaims{
		// 	ExpiresAt: int64(standardClaims["exp"].(float64)),
		// },
	}
}
