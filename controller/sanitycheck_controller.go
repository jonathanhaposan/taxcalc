package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jonathanhaposan/taxcalc/model"
	"github.com/jonathanhaposan/taxcalc/util"
)

type tmplAlert struct {
	HasError bool
	Message  string
}

func HandleSanityCheckList(w http.ResponseWriter, r *http.Request) {
	type result struct {
		Bills              []model.Bill `json:"bills"`
		TotalAmount        string       `json:"total_amount"`
		TotalTax           string       `json:"total_tax"`
		TotalOriginalPrice string       `json:"total_original_price"`
	}

	bills, totalAmount, totalTax, totalOriginalPrice, err := model.GetAllBill()
	if err != nil {
		log.Println("[HandleSanityCheckList] get data error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := &result{bills, totalAmount, totalTax, totalOriginalPrice}

	tmpl := template.Must(template.ParseFiles("assets/template/list.html"))
	tmpl.Execute(w, res)
}

func HandleSanityCheckSubmit(w http.ResponseWriter, r *http.Request) {
	var alert tmplAlert
	tmpl := template.Must(template.ParseFiles("assets/template/form_submit.html"))

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("[HandleSanityCheckSubmit] parse form error: ", err)
			alert = tmplAlert{true, err.Error()}
			tmpl.Execute(w, alert)
			return
		}

		productName, taxCode, price, err := validateData(r.Form)
		if err != nil {
			log.Println("[HandleSanityCheckSubmit] validation error: ", err)
			alert = tmplAlert{true, err.Error()}
			tmpl.Execute(w, alert)
			return
		}
		taxAmount, totalAmount := util.CalculateTax(taxCode, price)

		bill := model.Bill{}
		bill.ProductName = productName
		bill.OriginalPrice = price
		bill.TaxAmount = taxAmount
		bill.TotalAmount = totalAmount
		bill.Tax.ID = taxCode

		err = bill.AddNewBill()
		if err != nil {
			log.Println("[HandleSanityCheckSubmit] add new bill error: ", err)
			alert = tmplAlert{true, err.Error()}
			tmpl.Execute(w, alert)
			return
		}

		http.Redirect(w, r, "/sanitycheck/list", http.StatusSeeOther)
	}

	if r.Method == "GET" {
		tmpl.Execute(w, alert)
	}
}
