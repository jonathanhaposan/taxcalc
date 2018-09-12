package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jonathanhaposan/taxcalc/model"
	"github.com/jonathanhaposan/taxcalc/util"
)

type jsonMsg struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func HandleGetAllBill(w http.ResponseWriter, r *http.Request) {
	type result struct {
		Bills              []model.Bill `json:"bills"`
		TotalAmount        string       `json:"total_amount"`
		TotalTax           string       `json:"total_tax"`
		TotalOriginalPrice string       `json:"total_original_price"`
	}

	bills, totalAmount, totalTax, totalOriginalPrice, err := model.GetAllBill()
	if err != nil {
		log.Println("[HandleGetAllBill] get data error: ", err)
		msg := jsonMsg{
			http.StatusInternalServerError,
			err,
		}
		resp, _ := json.Marshal(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	data := &result{bills, totalAmount, totalTax, totalOriginalPrice}

	msg := jsonMsg{
		http.StatusOK,
		data,
	}
	resp, err := json.Marshal(msg)
	if err != nil {
		log.Println("[HandleGetAllBill] json marshal error: ", err)
		msg = jsonMsg{
			http.StatusInternalServerError,
			err,
		}

		resp, _ = json.Marshal(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func HandleSubmitBill(w http.ResponseWriter, r *http.Request) {
	var (
		msg  jsonMsg
		resp []byte
		bill model.Bill
	)

	err := r.ParseForm()
	if err != nil {
		log.Println("[HandleSubmitBill] parse form error: ", err)
		msg = jsonMsg{
			http.StatusInternalServerError,
			err,
		}

		resp, _ = json.Marshal(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	productName, taxCode, price, err := validateData(r.Form)
	if err != nil {
		log.Println("[HandleSubmitBill] validation error: ", err)
		msg = jsonMsg{
			http.StatusInternalServerError,
			err,
		}

		resp, _ = json.Marshal(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	taxAmount, totalAmount := util.CalculateTax(taxCode, price)

	bill = model.Bill{}
	bill.ProductName = productName
	bill.OriginalPrice = price
	bill.TaxAmount = taxAmount
	bill.TotalAmount = totalAmount
	bill.Tax.ID = taxCode

	err = bill.AddNewBill()
	if err != nil {
		log.Println("[HandleSubmitBill] add new bill error: ", err)
		msg = jsonMsg{
			http.StatusInternalServerError,
			err,
		}

		resp, _ = json.Marshal(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	msg = jsonMsg{
		http.StatusOK,
		"Success add bill",
	}

	resp, err = json.Marshal(msg)
	if err != nil {
		log.Println("[HandleSubmitBill] marshal error: ", err)
		msg = jsonMsg{
			http.StatusInternalServerError,
			err,
		}

		resp, _ = json.Marshal(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func validateData(data url.Values) (productName string, taxCode, price int64, err error) {
	if len(data.Get("product_name")) == 0 {
		err = errors.New("product name should not empty")
		log.Println("[validateData] product name error:", err)
		return
	}
	productName = data.Get("product_name")

	taxCode, err = strconv.ParseInt(data.Get("tax_code"), 10, 64)
	if err != nil {
		log.Println("[validateData] parse tax code error:", err)
		return
	}

	if taxCode != model.TaxCodeFood && taxCode != model.TaxCodeTobacco && taxCode != model.TaxCodeEntertainment {
		err = errors.New("wrong tax code")
		log.Println("[validateData] tax code error:", err)
		return
	}

	price, err = strconv.ParseInt(data.Get("amount"), 10, 64)
	if err != nil {
		log.Println("[validateData] parse price error:", err)
		return
	}

	if price == 0 {
		err = errors.New("price should not empty")
		log.Println("[validateData] price error:", err)
		return
	}

	return
}
