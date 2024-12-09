package abcToken

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// type JWTService interface {
// }

type IToken interface {
	GenerateToken(id string, email string, role string) (string, error)
}

type token struct {
	secretKey string
	issuer    string // created by
	audience  string // use by
}

func NewToken(secretKey string, issuer string) IToken {
	if secretKey == "" {
		secretKey = "secret"
	}
	if issuer == "" {
		issuer = "http://localhost:8080"
	}

	return &token{
		secretKey: secretKey,
		issuer:    issuer,
		audience:  issuer,
	}
}

type customClaims struct {
	// Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (builder *token) GenerateToken(id string, email string, role string) (string, error) {
	claims := customClaims{
		Email: email,
		Role:  role,
	}

	// assign standard claim
	claims.ExpiresAt = time.Now().Add(time.Hour * 48).Unix() // add 48 hour
	claims.Issuer = builder.issuer
	claims.Audience = builder.audience

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	tokenString, err := token.SignedString([]byte(builder.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
