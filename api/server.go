package api

import (
	"net/http"
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

var Service *bank.Service

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/bundesbank", index)
	r.HandleFunc("/bundesbank/v1/banks", bank.MakeBanksEndpoint(Service))
	r.HandleFunc("/bundesbank/v1/banks/{id}", bank.MakeBankEndpoint(Service))
	return r
}

func StartService() {
	log.Printf("Visit http://%s:%d/bundesbank for API docs...", getHostname(), port)
	http.ListenAndServe(":"+strconv.Itoa(port), Router())
}

type ResponseWrapper struct {
	Banks []bank.Bank
}

func index(w http.ResponseWriter, r *http.Request) {
	hostname := getHostname()
	index := Index{Hostname: hostname, Port: port}
	t, _ := template.New("index", Asset).Parse("api/index.html")
	t.Execute(w, struct{ Index }{index})
}

func getHostname() string {
	if hostname, err := os.Hostname(); err == nil && (hostname == "golem" || hostname == "Ralfs-MBP") {
		return "localhost"
	}

	return "94.130.79.196"
}
