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

func Test_validateData(t *testing.T) {
	var args1, args2, args3, args4, args5, args6 url.Values

	args1 = make(url.Values)
	args2 = make(url.Values)
	args3 = make(url.Values)
	args4 = make(url.Values)
	args5 = make(url.Values)
	args6 = make(url.Values)

	args1.Add("product_name", "")
	args2.Add("product_name", "kitkat")
	args2.Add("tax_code", "")
	args3.Add("product_name", "kitkat")
	args3.Add("tax_code", "1123")
	args4.Add("product_name", "kitkat")
	args4.Add("tax_code", "1")
	args4.Add("amount", "")
	args5.Add("product_name", "kitkat")
	args5.Add("tax_code", "1")
	args5.Add("amount", "0")
	args6.Add("product_name", "kitkat")
	args6.Add("tax_code", "1")
	args6.Add("amount", "1000")

	type args struct {
		data url.Values
	}
	tests := []struct {
		name            string
		args            args
		wantProductName string
		wantTaxCode     int64
		wantPrice       int64
		wantErr         bool
	}{
		{"1", args{args1}, "", 0, 0, true},
		{"2", args{args2}, "kitkat", 0, 0, true},
		{"3", args{args3}, "kitkat", 1123, 0, true},
		{"4", args{args4}, "kitkat", 1, 0, true},
		{"5", args{args5}, "kitkat", 1, 0, true},
		{"6", args{args6}, "kitkat", 1, 1000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProductName, gotTaxCode, gotPrice, err := validateData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotProductName != tt.wantProductName {
				t.Errorf("validateData() gotProductName = %v, want %v", gotProductName, tt.wantProductName)
			}
			if gotTaxCode != tt.wantTaxCode {
				t.Errorf("validateData() gotTaxCode = %v, want %v", gotTaxCode, tt.wantTaxCode)
			}
			if gotPrice != tt.wantPrice {
				t.Errorf("validateData() gotPrice = %v, want %v", gotPrice, tt.wantPrice)
			}
		})
	}
}
