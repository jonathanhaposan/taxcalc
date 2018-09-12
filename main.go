package main

import (
	"log"
	"net/http"

	"github.com/jonathanhaposan/taxcalc/database"
	"github.com/jonathanhaposan/taxcalc/model"
	"github.com/jonathanhaposan/taxcalc/router"
)

const (
	Test123 = 123
)

func main() {
	r := router.Init()
	db := database.Init()
	model.Init(db)

	log.Println("Server start on :9001")
	http.ListenAndServe(":9001", r)
}
