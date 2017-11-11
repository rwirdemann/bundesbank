package service

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
	"bitbucket.org/rwirdemann/bundesbank/util"
	"github.com/arschles/go-bindata-html-template"
	"bitbucket.org/rwirdemann/bundesbank/html"
)

// Data struct for index.html
type Index struct {
	Hostname string
	Port     int
}

const port = 8091

func StartService() {
	r := http.NewServeMux()
	r.HandleFunc("/bundesbank", index)
	r.HandleFunc("/bundesbank/v1/banks", banks)

	hostname := util.GetHostname()
	log.Printf("Visit http://%s:%d/bundesbank for API docs...", hostname, port)
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}

type ResponseWrapper struct {
	Banks []Bank
}

func banks(w http.ResponseWriter, r *http.Request) {
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

func queryByBlz(blz string, w http.ResponseWriter) {
	if banks, ok := BanksByPlz[blz]; ok {
		response := ResponseWrapper{Banks: banks}
		json := util.Json(response)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, json)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func queryByBic(bic string, w http.ResponseWriter) {
	if bankSlice, ok := BanksByBic[bic]; ok {
		response := ResponseWrapper{Banks: bankSlice}
		json := util.Json(response)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, json)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func queryByName(name string, w http.ResponseWriter) {
	if bankSlice, ok := BanksByBezeichnung[name]; ok {
		log.Printf("returning %d matchtes", len(bankSlice))
		response := ResponseWrapper{Banks: bankSlice}
		json := util.Json(response)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, json)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}


func index(w http.ResponseWriter, r *http.Request) {
	hostname := util.GetHostname()
	index := Index{Hostname: hostname, Port: port}
	template, _ := template.New("index", html.Asset).Parse("html/index.html")
	template.Execute(w, struct{ Index }{index})
}
