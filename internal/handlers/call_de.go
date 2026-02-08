package handlers

import (
	"encoding/json"
	"net/http"

	"duty-alert-func/internal/services"
)

type CallDutyRequest struct {
	Message string `json:"message"`
}

func CallDutyEngineer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req CallDutyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dutyService := services.NewDutyService()

	if err := dutyService.CallOnDutyEngineers(req.Message); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "calls triggered",
	})
}

