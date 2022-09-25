package handlers

import (
	"encoding/json"
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

type CurrencyData struct {
	Base           string
	LastUpdate     int64
	Exchange_rates map[string]float64
}

func (h handler) ShowCurrency(w http.ResponseWriter, r *http.Request) {

	currencies, err := currency.GetCurrencies(h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, currencies)
}

func (h handler) InsertCurrencyData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var currencyData CurrencyData

	err = json.Unmarshal(body, &currencyData)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	id, err := currency.GetCurrencyId(currencyData.Base, h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var currencies_data []models.CurrencyData

	for key, element := range currencyData.Exchange_rates {
		currencies_data = append(currencies_data, models.CurrencyData{
			Currency_ID: id,
			Base:        key,
			Rate:        element,
		})
	}

	if err := currency.InsertCurrencyData(h.DB, currencies_data); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, currencies_data)
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

func (h handler) InsertAllCurrencyData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var currencyData CurrencyData

	err = json.Unmarshal(body, &currencyData)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	id, err := currency.GetCurrencyId(currencyData.Base, h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var currencies_data []models.CurrencyData

	for key, element := range currencyData.Exchange_rates {
		currencies_data = append(currencies_data, models.CurrencyData{
			Currency_ID: id,
			Base:        key,
			Rate:        element,
		})
	}

	if err := currency.InsertCurrencyData(h.DB, currencies_data); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	for key, element := range currencyData.Exchange_rates {
		var rest_currencies_data []models.CurrencyData
		id, _ := currency.GetCurrencyId(key, h.DB)

		if id != 0 {
			for r_key, r_element := range currencyData.Exchange_rates {
				if key != r_key {
					new_rate := r_element / element
					rest_currencies_data = append(rest_currencies_data, models.CurrencyData{
						Currency_ID: id,
						Base:        r_key,
						Rate:        new_rate,
					})
				}
			}
			if err := currency.InsertCurrencyData(h.DB, rest_currencies_data); err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
		}
		rest_currencies_data = nil
	}

	responses.JSON(w, http.StatusCreated, currencies_data)
}
