package util

import (
	"github.com/jonathanhaposan/taxcalc/model"

	"github.com/shopspring/decimal"
)

func CalculateTax(taxCode, price int64) (float64, float64) {
	var (
		taxAmount   float64
		totalAmount float64
		tempTax     decimal.Decimal
		tempTot     decimal.Decimal
	)
	switch taxCode {
	case model.TaxCodeFood:
		tempTax = decimal.NewFromFloat(0.1).Mul(decimal.New(price, 0))
	case model.TaxCodeTobacco:
		tempTax = decimal.NewFromFloat(0.02).Mul(decimal.New(price, 0))
		tempTax = tempTax.Add(decimal.New(10, 0))
	case model.TaxCodeEntertainment:
		if price >= 100 {
			newPrice := price - 100
			tempTax = decimal.NewFromFloat(0.01).Mul(decimal.New(newPrice, 0))
		}
	}

	tempTot = tempTax.Add(decimal.New(price, 0))

	taxAmount, _ = tempTax.Float64()
	totalAmount, _ = tempTot.Float64()

	return taxAmount, totalAmount
}
