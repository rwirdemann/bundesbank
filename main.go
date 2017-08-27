package main

import (
	"fmt"
	"bitbucket.org/rwirdemann/bundesbank/service"
	"flag"
	"bitbucket.org/rwirdemann/bundesbank/util"
)

func main() {
	bundesbankFile := flag.String("f", "", "a blz file")
	flag.Parse()
	if *bundesbankFile == "" {
		fmt.Printf("usage: bundesbank -f 'blz-file.txt'\n")
		return
	}
	ip := util.GetHostname()
	fmt.Printf("IP: %v\n", ip)
	service.ImportBundesbankFile(*bundesbankFile)
	service.StartService()
}
