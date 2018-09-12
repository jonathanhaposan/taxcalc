package model

import (
	"fmt"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestBillAddNewBill(t *testing.T) {
	testBillAddNewBillPositive(t)
	testBillAddNewBillNegative(t)
}

func TestBillGetAllBill(t *testing.T) {
	testBillGetAllBillErrorQuery(t)
	testBillGetAllBillErrorScan(t)
	testBillGetAllBillPositive(t)
}

func testBillAddNewBillPositive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer db.Close()
	Init(db)

	bill := &Bill{
		ProductName:   "prodName",
		Tax:           Tax{ID: 123},
		OriginalPrice: 1000,
		TaxAmount:     10,
		TotalAmount:   1010,
	}

	mock.ExpectExec("INSERT INTO bill").WithArgs("prodName", 123, 1000, float64(10), float64(1010)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err = bill.AddNewBill(); err != nil {
		t.Errorf("Error was not expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}

func testBillAddNewBillNegative(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer db.Close()

	Init(db)

	bill := &Bill{
		ProductName:   "prodName",
		Tax:           Tax{ID: 123},
		OriginalPrice: 1000,
		TaxAmount:     10,
		TotalAmount:   1010,
	}

	mock.ExpectExec("INSERT INTO bill").WithArgs("prodName", 123, 1000, float64(10), float64(1010)).
		WillReturnError(fmt.Errorf("Error"))

	if err = bill.AddNewBill(); err == nil {
		t.Errorf("Error was expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}

func testBillGetAllBillErrorQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer db.Close()
	Init(db)

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnError(fmt.Errorf("Error"))

	_, _, _, _, err = GetAllBill()

	if err == nil {
		t.Errorf("Error was expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}

func testBillGetAllBillErrorScan(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer db.Close()
	Init(db)

	rows := sqlmock.NewRows([]string{"bill_id", "product_name", "tax_code_id", "description", "original_price", "tax_amount", "total_amount"}).
		AddRow("1", "Disc", "trouble maker", "Food", 12312, 123.12, 1231231)

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	_, _, _, _, err = GetAllBill()

	if err == nil {
		t.Errorf("Error was expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}

func testBillGetAllBillPositive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer db.Close()
	Init(db)

	rows := sqlmock.NewRows([]string{"bill_id", "product_name", "tax_code_id", "description", "original_price", "tax_amount", "total_amount"}).
		AddRow(1, "Disc", 1, "Food", 12312, 123.12, 1231231)

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	_, _, _, _, err = GetAllBill()

	if err != nil {
		t.Errorf("Error was expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}
