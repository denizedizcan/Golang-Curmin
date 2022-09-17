package handlers

import (
	"net/http"

	"github.com/denizedizcan/Golang-Curmin/common/responses"
	"github.com/denizedizcan/Golang-Curmin/pkg/currency"
)

func (h handler) ShowCurrency(w http.ResponseWriter, r *http.Request) {

	currencies, err := currency.GetCurrencies(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, currencies)
}
