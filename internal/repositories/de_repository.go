package repositories

import "duty-alert-func/internal/models"

type DutyRepository interface {
	GetOnDutyEngineers() ([]models.DutyEngineer, error)
}
