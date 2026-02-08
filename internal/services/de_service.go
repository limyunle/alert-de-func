package services

import (
	"fmt"
)

type DutyService struct {
	excelService   *ExcelService
	infobipService *InfobipService
}

func NewDutyService() *DutyService {
	return &DutyService{
		excelService:   NewExcelService(),
		infobipService: NewInfobipService(),
	}
}

func (d *DutyService) CallOnDutyEngineers(alertMessage string) error {
	engineers, err := d.excelService.GetOnDutyEngineers()
	if err != nil {
		return err
	}

	for _, de := range engineers {
		msg := fmt.Sprintf(
			"Hello %s. You are on duty. Alert message: %s",
			de.Name,
			alertMessage,
		)

		if err := d.infobipService.Call(de.Mobile, msg); err != nil {
			return err
		}
	}

	return nil
}

