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

func marshal(entities interface{}) string {
	if b, err := json.Marshal(entities); err == nil {
		return fmt.Sprintf("%s", string(b[:]))
	}
	return ""
}

