package dal

//mockgen -destination=internal/dal/mock/mockrepo.go github.com/bgoldovsky/kaspersky-fan-in/internal/dal Repo

import models "github.com/bgoldovsky/counter-api/internal/models"

// Repo iface
type Repo interface {
	Save(s *models.Counter) error
	Get() (*models.Counter, error)
}
