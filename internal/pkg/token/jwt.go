package token

import (
	"errors"
	"time"

	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/golang-jwt/jwt/v4"
)

// minSecretKeySize is the minimum size of the secret key used for signing tokens.
const minSecretKeySize = 32

// Jwt is a JSON Web Token implementation.
type Jwt struct {
	secretKey string
	duration  time.Duration
}

// newJWT creates a new JWT token generator.
func newJWT(secretKey string, duration time.Duration) (*Jwt, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, helper.ErrInvalidToken
	}

	return &Jwt{
		secretKey,
		duration,
	}, nil
}

// Create creates a new token for a specific email and duration.
func (j *Jwt) Create(email string) (string, error) {
	payload, err := NewPayload(email, j.duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(j.secretKey))

	return token, err
}

// Verify checks if the token is valid or not
func (j *Jwt) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, helper.ErrInvalidToken
		}

		return []byte(j.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationErr.Inner, helper.ErrExpiredToken) {
			return nil, helper.ErrExpiredToken
		}

		return nil, helper.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, helper.ErrInvalidToken
	}

	return payload, nil
}
