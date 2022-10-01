package reshape

import "github.com/denizedizcan/Golang-Curmin/pkg/currency"

type CurrencyData struct {
	Currencies []currency.CurrencyShow `json:"currencies"`
}

func ReshapeCurrencyList(currencylist []currency.CurrencyShow) CurrencyData {
	var currencyData CurrencyData
	currencyData.Currencies = append(currencyData.Currencies, currencylist...)
	return currencyData
}
