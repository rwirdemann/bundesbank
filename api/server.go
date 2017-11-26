package api

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
	"github.com/arschles/go-bindata-html-template"
	"bitbucket.org/rwirdemann/bundesbank/bank"
	"github.com/gorilla/mux"
	"os"
)

// Data struct for index.html
type Index struct {
	Hostname string
	Port     int
}

const port = 8091

var Repository bank.BankRepository

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/bundesbank", index)
	r.HandleFunc("/bundesbank/v1/banks", banksHandler)
	r.HandleFunc("/bundesbank/v1/banks/{id}", bankHandler)
	return r
}

func StartService() {
	log.Printf("Visit http://%s:%d/bundesbank for API docs...", getHostname(), port)
	http.ListenAndServe(":"+strconv.Itoa(port), Router())
}

func bankHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if idParam, ok := vars["id"]; ok {
		if id, err := strconv.Atoi(idParam); err == nil {
			if bank, ok := bank.GetRepositoryInstance().ById(id); ok {
				json := Json(bank)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, json)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

type ResponseWrapper struct {
	Banks []bank.Bank
}

func banksHandler(w http.ResponseWriter, r *http.Request) {
	if blz, ok := r.URL.Query()["blz"]; ok {
		queryByBlz(blz[0], w)
	}

	if bic, ok := r.URL.Query()["bic"]; ok {
		queryByBic(bic[0], w)
	}

	if name, ok := r.URL.Query()["name"]; ok {
		queryByName(name[0], w)
	}
}

func writeResponse(banks []bank.Bank, w http.ResponseWriter) {
	response := ResponseWrapper{Banks: banks}
	json := Json(response)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, json)
}

func queryByBlz(blz string, w http.ResponseWriter) {
	if banks, ok := Repository.ByBlz(blz); ok {
		writeResponse(banks, w)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func queryByBic(bic string, w http.ResponseWriter) {
	if banks, ok := Repository.ByBic(bic); ok {
		writeResponse(banks, w)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func queryByName(name string, w http.ResponseWriter) {
	if banks, ok := Repository.ByBezeichnung(name); ok {
		writeResponse(banks, w)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	hostname := getHostname()
	index := Index{Hostname: hostname, Port: port}
	t, _ := template.New("index", Asset).Parse("api/index.html")
	t.Execute(w, struct{ Index }{index})
}

func getHostname() string {
	if hostname, err := os.Hostname(); err == nil && hostname == "golem" {
		return "localhost"
	}

	return "94.130.79.196"
}
