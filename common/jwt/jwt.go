package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	secretKey     string
	tokenDuration time.Duration
}

type CustomClaims struct {
	StandardClaims jwt.StandardClaims
	Nim            string `json:"nim"`
	Role           uint32 `json:"role"`
}

func NewJWT(secretKey string, tokenDuration time.Duration) *JWT {
	return &JWT{
		secretKey,
		tokenDuration,
	}
}

func (j *JWT) GenerateToken(nim string, role uint32) (string, error) {
	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(j.tokenDuration).Unix(),
		},
		Nim:  nim,
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWT) Verify(accessToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if err := claims.Valid(); err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}

func (c *CustomClaims) Valid() error {
	// You need to implement your own validation logic here.
	// For example, you can check if the token has expired.
	if time.Now().Unix() > c.StandardClaims.ExpiresAt {
		return fmt.Errorf("token has expired")
	}

	// Add additional validation checks as needed.

	return nil
}
