package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/jonathanhaposan/taxcalc/model"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestHandleGetAllBill(t *testing.T) {
	type result struct {
		Bills              []model.Bill `json:"bills"`
		TotalAmount        string       `json:"total_amount"`
		TotalTax           string       `json:"total_tax"`
		TotalOriginalPrice string       `json:"total_original_price"`
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer db.Close()
	model.Init(db)

	rows := sqlmock.NewRows([]string{"product_name", "tax_code_id", "description", "original_price", "tax_amount", "total_amount"}).
		AddRow("test", 1, "Food", 1000, 100, 1100)

	bill := []model.Bill{
		{
			ProductName:   "test",
			Tax:           model.Tax{ID: 1, Type: "Food"},
			OriginalPrice: 1000,
			TaxAmount:     float64(100),
			TotalAmount:   float64(1100),
		},
	}

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(HandleGetAllBill)
	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("return wrong status: got %+v want %+v", status, http.StatusOK)
	}

	expected := jsonMsg{http.StatusOK, result{bill, "1100", "100", "1000"}}
	actual := recorder.Body.Bytes()

	expectedByte, _ := json.Marshal(expected)

	if bytes.Compare(expectedByte, actual) != 0 {
		t.Errorf("return unexpected body: got %+v want %+v", recorder.Body.String(), string(expectedByte))
	}

}

func TestHandleSubmitBill(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer db.Close()
	model.Init(db)

	mock.ExpectExec("INSERT INTO bill").WithArgs("prodName", 1, 1000, float64(100), float64(1100)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	form := newCreateBillForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(HandleSubmitBill)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("return wrong status: got %+v want %+v", status, http.StatusOK)
	}

	expected := jsonMsg{http.StatusOK, "Success add bill"}
	actual := recorder.Body.Bytes()

	expectedByte, _ := json.Marshal(expected)

	if bytes.Compare(expectedByte, actual) != 0 {
		t.Errorf("return unexpected body: got %+v want %+v", recorder.Body.String(), string(expectedByte))
	}
}

func newCreateBillForm() *url.Values {
	form := url.Values{}
	form.Set("product_name", "prodName")
	form.Set("tax_code", "1")
	form.Set("amount", "1000")
	return &form
}
