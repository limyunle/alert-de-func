package services

import (
	"duty-alert-func/internal/models"
)

type ExcelService struct {
	// later:
	// graphClient *graph.Client
}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (e *ExcelService) GetOnDutyEngineers() ([]models.DutyEngineer, error) {
	// TODO:
	// - Call Graph
	// - Read OnDuty table
	// - Filter by current week

	return []models.DutyEngineer{
		{
			Name:        "Bobby",
			Mobile: "+6592479152",
		},
	}, nil
}
