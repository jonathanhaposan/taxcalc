package model

import (
	"log"
	"strconv"

	"github.com/shopspring/decimal"
)

type Bill struct {
	ID            int64   `json:"bill_id"`
	ProductName   string  `json:"product_name"`
	Tax           Tax     `json:"tax"`
	OriginalPrice int64   `json:"original_price"`
	TaxAmount     float64 `json:"tax_amount"`
	TotalAmount   float64 `json:"total_amount"`
}

func (b *Bill) AddNewBill() error {
	_, err := db.Exec(`INSERT INTO bill(product_name,tax_code_id,original_price, tax_amount, total_amount)
	VALUES($1,$2,$3,$4,$5)`, b.ProductName, b.Tax.ID, b.OriginalPrice, b.TaxAmount, b.TotalAmount)
	if err != nil {
		log.Println("[AddNewBill] error:", err)
		return err
	}

	return nil
}

func GetAllBill() ([]Bill, string, string, string, error) {
	var (
		bills              []Bill
		temp               Bill
		totalAmount        decimal.Decimal
		totalTaxAmount     decimal.Decimal
		totalOriginalPrice int64
	)

	query := `
	SELECT
		b.product_name,
		b.tax_code_id,
		tc.description,
		b.original_price,
		b.tax_amount,
		b.total_amount
	FROM
		bill b JOIN tax_code tc ON b.tax_code_id = tc.code_id
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Println("[AddNewBill] fetching data error:", err)
		return nil, "", "", "", err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&temp.ProductName, &temp.Tax.ID, &temp.Tax.Type, &temp.OriginalPrice, &temp.TaxAmount, &temp.TotalAmount)
		if err != nil {
			log.Println("[AddNewBill] scan query error:")
			return nil, "", "", "", err
		}

		totalAmount = totalAmount.Add(decimal.NewFromFloat(temp.TotalAmount))
		totalTaxAmount = totalTaxAmount.Add(decimal.NewFromFloat(temp.TaxAmount))
		totalOriginalPrice += temp.OriginalPrice
		bills = append(bills, temp)
	}

	strTotalOri := strconv.FormatInt(totalOriginalPrice, 10)

	return bills, totalAmount.String(), totalTaxAmount.String(), strTotalOri, nil

}
