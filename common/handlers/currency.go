package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/denizedizcan/Golang-Curmin/common/models"
	"github.com/denizedizcan/Golang-Curmin/common/reshape"
	"github.com/denizedizcan/Golang-Curmin/common/responses"
	"github.com/denizedizcan/Golang-Curmin/pkg/currency"
)

const layout = "2006-01-02T15:04:05"

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

	responses.JSON(w, http.StatusCreated, reshape.ReshapeCurrencyList(currencies))
}

func (h handler) GetTimeCurrency(w http.ResponseWriter, r *http.Request) {
	base_list := r.URL.Query()["base"]
	target_list := r.URL.Query()["target"]
	start_date_list := r.URL.Query()["start_date"]
	end_date_list := r.URL.Query()["end_date"]
	var (
		base       string
		target     string
		start_date string
		end_date   string
		date_gt    time.Time
		date_lt    time.Time
	)
	loc, _ := time.LoadLocation("Europe/Istanbul")
	if len(target_list) > 0 {
		target = target_list[0]
	}
	if len(base_list) > 0 {
		base = base_list[0]
	}
	if len(start_date_list) > 0 {
		start_date = start_date_list[0]
		date_gt, _ = time.ParseInLocation(layout, start_date, loc)
	} else {
		date_gt, _ = time.ParseInLocation(layout, "1970-01-02T15:04:05", loc)
	}
	if len(end_date_list) > 0 {
		end_date = end_date_list[0]
		date_lt, _ = time.ParseInLocation(layout, end_date, loc)
	} else {
		date_lt = time.Now()
		date_lt, _ = time.ParseInLocation(date_lt.Format(layout), layout, loc)
	}
	if base == "" {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Required base parameter"))
		return
	}
	if target == "" {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Required target parameter"))
		return
	}
	currency, err := currency.GetCurrencyBytime(base, target, date_gt, date_lt, h.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, currency)
}

func (h handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	base_list := r.URL.Query()["base"]
	target_list := r.URL.Query()["target"]
	var (
		base   string
		target string
	)
	if len(target_list) > 0 {
		target = target_list[0]
	}
	if len(base_list) > 0 {
		base = base_list[0]
	}
	if base == "" {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Required base parameter"))
		return
	}
	if target != "" {
		currency, err := currency.GetCurrencyByBaseAndTarget(base, target, h.DB)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		responses.JSON(w, http.StatusAccepted, currency)

	} else {
		currency, err := currency.GetCurrencyByBase(base, h.DB)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		responses.JSON(w, http.StatusAccepted, currency)
	}

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
