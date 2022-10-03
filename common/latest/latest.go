package latest

import (
	"sort"

	"github.com/denizedizcan/Golang-Curmin/common/models"
)

func FindLatestTargetedData(currencies_data models.Currency) models.Currency {
	sort.Slice(currencies_data.CurrencyData, func(i, j int) bool {
		return currencies_data.CurrencyData[i].Date.After(currencies_data.CurrencyData[j].Date)
	})
	currencies_data.CurrencyData = currencies_data.CurrencyData[0:1]
	return currencies_data
}

func FindLatestData(currencies_data models.Currency) models.Currency {
	max_date := currencies_data.CurrencyData[0].Date
	for _, items := range currencies_data.CurrencyData {

		if max_date.Before(items.Date) {
			max_date = items.Date
		}
	}

	var new_data_list []models.CurrencyData
	for _, items := range currencies_data.CurrencyData {

		if items.Date == max_date {
			new_data_list = append(new_data_list, items)
		}
	}
	currencies_data.CurrencyData = new_data_list
	return currencies_data
}
