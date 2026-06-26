package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	DefaultTokenDuration = 24 * 7 * 60 * 60 // 7 days in seconds
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userId uint, email string, name string, role string) (string, error)
	// ValidateToken(tokenStr string) (*JWTClaims, error)
}
type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey:     secretKey,
		tokenDuration: DefaultTokenDuration,
	}
}

func (js *jwtService) GenerateToken(userId uint, email string, name string, role string) (string, error) {

	// create claims
	claims := JWTClaims{
		UserID: userId,
		Name:   name,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(js.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "spotsync",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // create token with claims

	tokenString, err := token.SignedString([]byte(js.secretKey)) // sign token with secret key
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
