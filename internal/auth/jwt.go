package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	DefaultTokenDuration = 7 * 24 * time.Hour
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userId uint, email string, name string, role string) (string, error)
	ValidateToken(tokenStr string) (*JWTClaims, error)
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

func (js *jwtService) ValidateToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(js.secretKey), nil

	})

	if err != nil {
		return nil, fmt.Errorf("Unexpected signing method: %w", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
