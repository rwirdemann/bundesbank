package main

import (
	"fmt"
	"bitbucket.org/rwirdemann/bundesbank/api"
	"flag"
	"bitbucket.org/rwirdemann/bundesbank/bank"
	"bitbucket.org/rwirdemann/bundesbank/parser"
)

func main() {
	bundesbankFile := flag.String("f", "", "a blz file")
	flag.Parse()
	if *bundesbankFile == "" {
		fmt.Printf("usage: bundesbank -f 'blz-file.txt'\n")
		return
	}
	parser.ImportBundesbankFile(*bundesbankFile)
	api.Repository = bank.GetRepositoryInstance()
	api.StartService()
}
