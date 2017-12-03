package api

import (
	"bitbucket.org/rwirdemann/bundesbank/bank"
	"github.com/gorilla/mux"
)

// Data struct for index.html
type Index struct {
	Hostname string
	Port     int
}

var Service *bank.Service

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/bundesbank/v1/banks", bank.MakeBanksEndpoint(Service))
	r.HandleFunc("/bundesbank/v1/banks/{id}", bank.MakeBankEndpoint(Service))
	return r
}

func StartService() {
}
