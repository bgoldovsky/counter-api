package dal

import models "github.com/bgoldovsky/counter-api/internal/models"

// Repo iface
type Repo interface {
	Save(s *models.Counter) error
	Get() (*models.Counter, error)
}
