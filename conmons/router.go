package conmons

import "github.com/gorilla/mux"

func SetCodeRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/valor", CodeRest).Methods("POST")
	subRoute.HandleFunc("/passphrase", Passphrase).Methods("POST")
}
