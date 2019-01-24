package domains

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Base base model for domains
type Base struct {
	ID        uuid.UUID  `sql:",type:uuid" json:"id"`
	CreatedAt time.Time  `sql:"default:now()" json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
