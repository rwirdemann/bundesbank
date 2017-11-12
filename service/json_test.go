package service

import (
	"testing"
	"bitbucket.org/rwirdemann/bundesbank/util"
	"bitbucket.org/rwirdemann/bundesbank/domain"
)

func TestSerializeBankResponse(t *testing.T) {
	response := ResponseWrapper{Banks: []domain.Bank{{Blz: "12345"}}}
	json := util.Json(response)
	expected := `{"Banks":[{"Id":0,"Blz":"12345","Bankleitzahlfuehrend":"","Bezeichnung":"","PLZ":"","Kurzbezeichnung":"","Pan":"","BIC":"","Pruefzifferberechnungsmethode":"","Datensatznummer":"","Aenderungskennzeichen":"","Bankleitzahlloeschung":"","Nachfolgebankleitzahl":""}]}`
	util.AssertEquals(t, expected, json)
}
