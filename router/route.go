package router

import (
	"net/http"

	"github.com/jonathanhaposan/taxcalc/controller"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	registerRouter(router)
	registerStaticFile(router)

	return router
}

func registerRouter(r *mux.Router) {
	r.HandleFunc("/bill/list", controller.HandleGetAllBill).Methods("GET")
	r.HandleFunc("/bill", controller.HandleSubmitBill).Methods("POST")

	r.HandleFunc("/sanitycheck/list", controller.HandleSanityCheckList).Methods("GET")
	r.HandleFunc("/sanitycheck/submit", controller.HandleSanityCheckSubmit).Methods("GET")
	r.HandleFunc("/sanitycheck/submit", controller.HandleSanityCheckSubmit).Methods("POST")
}

func registerStaticFile(r *mux.Router) {
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
}
