package router

import (
	"mongoserver/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Homepage).Methods("GET")

	router.HandleFunc("/locations", controller.StoreLocation).Methods("POST")
	router.HandleFunc("/locations/{category}", controller.GetLocationByCategory).Methods("GET")
	router.HandleFunc("/search", controller.SearchLocations).Methods("GET")
	router.HandleFunc("/trip-cost/{location_id}", controller.FetchCost).Methods("GET")
	return router
}
