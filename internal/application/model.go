package application

import (
	"time"

	"github.com/google/uuid"
)

type Connector struct {
	Provider     string
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

type Application struct {
	ID         uuid.UUID
	OwnerID    uuid.UUID
	Name       string
	Connectors []Connector
	CreatedAt  time.Time
}

type ApplicationDTO struct {
	ID         string      `json:"id"`
	OwnerID    string      `json:"owner_id"`
	Name       string      `json:"name"`
	Connectors []Connector `json:"connectors"`
	CreatedAt  string      `json:"created_at"`
}

// Converts an Application to a DTO for the API
func (a *Application) ToDTO() *ApplicationDTO {
	return &ApplicationDTO{
		ID:         a.ID.String(),
		OwnerID:    a.OwnerID.String(),
		Name:       a.Name,
		Connectors: a.Connectors,
		CreatedAt:  a.CreatedAt.Format(time.UnixDate),
	}
}
