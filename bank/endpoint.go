package bank

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"encoding/json"
)

func MakeBankEndpoint(service *Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if idParam, ok := vars["id"]; ok {
			if id, err := strconv.Atoi(idParam); err == nil {
				if bank, ok := service.byId(id); ok {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintf(w, marshal(bank))
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	}
}

func MakeBanksEndpoint(s *Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if blz, ok := r.URL.Query()["blz"]; ok {
			queryByBlz(blz[0], s, w)
		}

		if bic, ok := r.URL.Query()["bic"]; ok {
			queryByBic(bic[0], s, w)
		}

		if name, ok := r.URL.Query()["name"]; ok {
			queryByName(name[0], s, w)
		}
	}
}

type ResponseWrapper struct {
	Banks []Bank
}

func writeResponse(banks []Bank, w http.ResponseWriter) {
	response := ResponseWrapper{Banks: banks}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, marshal(response))
}

func queryByBlz(blz string, s *Service, w http.ResponseWriter) {
	if banks, ok := s.Repository.ByBlz(blz); ok {
		writeResponse(banks, w)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func queryByBic(bic string, s *Service, w http.ResponseWriter) {
	if banks, ok := s.Repository.ByBic(bic); ok {
		writeResponse(banks, w)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func queryByName(name string, s *Service, w http.ResponseWriter) {
	if banks, ok := s.Repository.ByBezeichnung(name); ok {
		writeResponse(banks, w)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func marshal(entities interface{}) string {
	if b, err := json.Marshal(entities); err == nil {
		return fmt.Sprintf("%s", string(b[:]))
	}
	return ""
}
