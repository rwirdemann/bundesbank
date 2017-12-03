package main

import (
	"flag"
	"fmt"

	"bitbucket.org/rwirdemann/bundesbank/bank"
	"bitbucket.org/rwirdemann/bundesbank/import"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

const port = 8091

func main() {

	// Handle command line args
	bundesbankFile := flag.String("f", "", "a blz file")
	flag.Parse()
	if *bundesbankFile == "" {
		fmt.Printf("usage: bundesbank -f 'blz-file.txt'\n")
		return
	}

	// Sertup services
	s := bank.NewBankService(bank.NewFileRepository())
	_import.ImportBundesbankFile(*bundesbankFile, s)

	// Register endpoints and start server
	r := mux.NewRouter()
	r.HandleFunc("/bundesbank/v1/banks", bank.MakeBanksEndpoint(s))
	r.HandleFunc("/bundesbank/v1/banks/{id}", bank.MakeBankEndpoint(s))

	fmt.Printf("http://%s:%d/bundesbank/v1/banks?blz=10010424", "localhost", port)
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
