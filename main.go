package main

import (
	"fmt"
	"bitbucket.org/rwirdemann/bundesbank/service"
	"flag"
	"bitbucket.org/rwirdemann/bundesbank/domain"
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
	service.Repository = domain.GetRepositoryInstance()
	service.StartService()
}
