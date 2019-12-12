package services

import "github.com/bgoldovsky/counter-api/internal/models"

// Counter iface
type Counter interface {
	Inc() error
	State() models.State
}
