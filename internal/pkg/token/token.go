package token

import (
	"github.com/bagashiz/gommerce/internal/pkg/config"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// Token is an interface for token implementations
type Token interface {
	// Create creates a new token for a specific user id, role, and duration
	Create(userID uint, isAdmin bool) (string, error)
	// Verify checks if the token is valid or not
	Verify(token string) (*Payload, error)
}

// New creates a new token instance based on the given token type
func New(tokenCfg *config.Token) (Token, error) {
	switch tokenCfg.Type {
	case "jwt":
		return newJWT(tokenCfg.SymmetricKey, tokenCfg.Duration)
	default:
		return nil, helper.ErrUnsupportedTokenType
	}
	// TODO: Add support for other token types
}
