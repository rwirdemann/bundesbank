package service

import (
	"testing"
	"bitbucket.org/rwirdemann/bundesbank/util"
)

func TestSerializeBankResponse(t *testing.T) {
	response := ResponseWrapper{Banks: []Bank{{Blz: "12345"}}}
	json := util.Json(response)
	expected := `{"Banks":[{"Blz":"12345","Bankleitzahlfuehrend":"","Bezeichnung":"","PLZ":"","Kurzbezeichnung":"","Pan":"","BIC":"","Pruefzifferberechnungsmethode":"","Datensatznummer":"","Aenderungskennzeichen":"","Bankleitzahlloeschung":"","Nachfolgebankleitzahl":""}]}`
	util.AssertEquals(t, expected, json)
}
