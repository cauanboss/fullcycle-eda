package controller

import (
	"encoding/json"
	"microservice/application"
	"microservice/domain/dto"
	"net/http"
	"strings"
)

type BalanceController struct {
	UseCases *application.BalanceUseCases
}

func NewBalanceController(useCases *application.BalanceUseCases) *BalanceController {
	return &BalanceController{
		UseCases: useCases,
	}
}

func (c *BalanceController) CreateBalance(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateBalanceInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	balance, err := c.UseCases.CreateBalance.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}

func (c *BalanceController) FindById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/balance/")

	balance, err := c.UseCases.FindById.Execute(&dto.FindByIdInput{
		ID: id,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}

func (c *BalanceController) Find(w http.ResponseWriter, r *http.Request) {
	balance, err := c.UseCases.Find.Execute(nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}
