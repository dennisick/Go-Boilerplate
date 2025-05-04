package model

import (
	"server/internal/domain/application/model"
	"time"

	"github.com/google/uuid"
)

type ConnectionProvider int

const (
	IMAP      ConnectionProvider = 1
	Google    ConnectionProvider = 2
	Microsoft ConnectionProvider = 3
)

type Connection struct {
	ID            uuid.UUID
	ApplicationID uuid.UUID
	Email         string
	Provider      ConnectionProvider
	RefreshToken  string
	CreatedAt     time.Time
	Application   *model.Application
}

type ConnectionDTO struct {
	ID            string `json:"id"`
	ApplicationID string `json:"application_id"`
	Email         string `json:"email"`
	Provider      string `json:"provider"`
	CreatedAt     string `json:"created_at"`
}

func (c *Connection) ToDTO() *ConnectionDTO {
	return &ConnectionDTO{
		ID:            c.ID.String(),
		ApplicationID: c.ApplicationID.String(),
		Email:         c.Email,
		Provider:      c.Provider.String(),
		CreatedAt:     c.CreatedAt.Format(time.UnixDate),
	}
}

func (p *ConnectionProvider) String() string {
	switch *p {
	case IMAP:
		return "IMAP"
	case Google:
		return "Google"
	case Microsoft:
		return "Microsoft"
	}

	return "unknown"
}
