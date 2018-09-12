package model

const (
	TaxCodeFood          = 1
	TaxCodeTobacco       = 2
	TaxCodeEntertainment = 3
)

type Tax struct {
	ID   int64  `json:"tax_id"`
	Type string `json:"type"`
}
