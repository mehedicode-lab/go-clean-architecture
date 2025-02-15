package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mehedicode-lab/go-clean-architecture/config"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Generate access & refresh tokens
func GenerateTokens(email string) (string, string, error) {
	// Access Token (valid for 15 minutes)
	accessClaims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.AppConfig.JwtConfig.AccessTTL) * time.Second)),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(config.AppConfig.JwtConfig.AccessSecret))
	if err != nil {
		return "", "", err
	}
	// Refresh Token (valid for 7 days)
	refreshClaims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.AppConfig.JwtConfig.RefreshTTL) * time.Second)),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(config.AppConfig.JwtConfig.RefreshSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Validate Access Token
func ValidateAccessToken(tokenString string) (*Claims, error) {
	return validateToken(tokenString, []byte(config.AppConfig.JwtConfig.AccessSecret))
}

// Validate Refresh Token
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	return validateToken(tokenString, []byte(config.AppConfig.JwtConfig.RefreshSecret))
}

// Common validation function
func validateToken(tokenString string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	// Checking for parse errors
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Checking if the token is not valid
	if !token.Valid {
		return nil, errors.New("token is invalid: token validation failed")
	}

	// Verifying the claims type and returning them
	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}

	// If the claims are of the wrong type
	return nil, errors.New("invalid token: claims are of the wrong type")
}
