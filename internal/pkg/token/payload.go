package token

import (
	"time"

	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/google/uuid"
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uint      `json:"user_id"`
	IsAdmin   bool      `json:"is_admin"`
	IsUsedAt  time.Time `json:"isused_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific user id, role, and duration
func NewPayload(userID uint, isAdmin bool, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IsAdmin:   isAdmin,
		IsUsedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return helper.ErrExpiredToken
	}

	return nil
}
