package handlers

import (
	"net/http"

	"github.com/denizedizcan/Golang-Curmin/models"
	"github.com/denizedizcan/Golang-Curmin/responses"
)

func (h handler) ShowCurrency(w http.ResponseWriter, r *http.Request) {

	currencies, err := models.FindAllCurrencys(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, currencies)
}
