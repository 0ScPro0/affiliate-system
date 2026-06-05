package core_security

import (
	"fmt"
	"time"

	"github.com/0ScPro0/affiliate-system/internal/core/config"
	"github.com/golang-jwt/jwt/v5"
)

// TokenType represents JWT token type (access or refresh).
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// CreateAccessToken generates a new access JWT with the given claims.
func CreateAccessToken(
	cfg *config.Config,
	sub map[string]any,
) (string, error) {
	token, err := createToken(cfg, sub, AccessToken)
	if err != nil {
		return "", fmt.Errorf("create access token: %w", err)
	}
	return token, nil
}

// CreateRefreshToken generates a new refresh JWT with the given claims.
func CreateRefreshToken(
	cfg *config.Config,
	sub map[string]any,
) (string, error) {
	token, err := createToken(cfg, sub, RefreshToken)
	if err != nil {
		return "", fmt.Errorf("create refresh token: %w", err)
	}
	return token, nil
}

// VerifyToken parses and validates a JWT token string, returns the raw claims.
func VerifyToken(
	cfg *config.Config,
	tokenString string,
) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Security.SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

// GetSub extracts the "sub" field from JWT claims as a generic map.
func GetSub(claims jwt.MapClaims) (map[string]any, error) {
	sub, ok := claims["sub"]
	if !ok {
		return nil, fmt.Errorf("missing 'sub' in token claims")
	}

	subMap, ok := sub.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("'sub' is not a map")
	}

	return subMap, nil
}

// createToken is the shared JWT creation logic.
func createToken(
	cfg *config.Config,
	sub map[string]any,
	tokenType TokenType,
) (string, error) {
	if !validateTokenType(tokenType) {
		return "", fmt.Errorf("unknown token type: %s", tokenType)
	}

	delta, err := calculateTokenExpireDelta(cfg, tokenType)
	if err != nil {
		return "", fmt.Errorf("calculate token expire delta: %w", err)
	}
	expire := time.Now().UTC().Add(delta)

	claims := jwt.MapClaims{
		"exp":  jwt.NewNumericDate(expire),
		"iat":  jwt.NewNumericDate(time.Now().UTC()),
		"type": string(tokenType),
		"iss":  "affiliate",
	}

	for key, value := range sub {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.Security.SecretKey))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return tokenString, nil
}

// validateTokenType checks if the token type is valid.
func validateTokenType(tokenType TokenType) bool {
	return tokenType == AccessToken || tokenType == RefreshToken
}

// calculateTokenExpireDelta returns token TTL from config.
func calculateTokenExpireDelta(
	cfg *config.Config,
	tokenType TokenType,
) (time.Duration, error) {
	switch tokenType {
	case AccessToken:
		return time.Duration(cfg.Security.AccessTokenExpireMinutes) * time.Minute, nil

	case RefreshToken:
		return time.Duration(cfg.Security.RefreshTokenExpireDays) * time.Hour * 24, nil

	default:
		return 0, fmt.Errorf("unknown token type: %s", tokenType)
	}
}