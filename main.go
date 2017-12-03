package main

import (
	"flag"
	"fmt"

	"bitbucket.org/rwirdemann/bundesbank/api"
	"bitbucket.org/rwirdemann/bundesbank/bank"
	"bitbucket.org/rwirdemann/bundesbank/import"
)

func main() {
	bundesbankFile := flag.String("f", "", "a blz file")
	flag.Parse()
	if *bundesbankFile == "" {
		fmt.Printf("usage: bundesbank -f 'blz-file.txt'\n")
		return
	}

	api.Service = bank.NewBankService(bank.NewFileRepository())
	_import.ImportBundesbankFile(*bundesbankFile, api.Service)
	api.StartService()
}
