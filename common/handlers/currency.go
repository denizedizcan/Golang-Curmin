package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/denizedizcan/Golang-Curmin/common/models"
	"github.com/denizedizcan/Golang-Curmin/common/responses"
	"github.com/denizedizcan/Golang-Curmin/pkg/currency"
)

type CurrencyList struct {
	Succes  bool
	Symbols map[string]string
}

func (h handler) ShowCurrency(w http.ResponseWriter, r *http.Request) {

	currencies, err := currency.GetCurrencies(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, currencies)
}

func (h handler) InsertCurrency(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var currencyList CurrencyList

	err = json.Unmarshal(body, &currencyList)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var currencies []models.Currency

	for key, element := range currencyList.Symbols {
		fmt.Println("Key:", key, "=>", "Element:", element)
		currencies = append(currencies, models.Currency{
			Code: key,
			Name: element,
		})
	}

	if err := currency.InsertCurrency(h.DB, currencies); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, currencyList)
}
