package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bitbucket.org/rwirdemann/bundesbank/bank"
)

var service *bank.Service

func init() {
	Service = bank.NewBankService(bank.NewFileRepository())
	bank.ImportBundesbankFile("service_test_data.txt", Service)
}

const b1 = `{"Id":1,"Blz":"10010424","Bankleitzahlfuehrend":"","Bezeichnung":"Aareal Bank","PLZ":"10666","Kurzbezeichnung":"Aareal Bank","Pan":"26910","BIC":"AARBDE5W100","Pruefzifferberechnungsmethode":"09","Datensatznummer":"004795","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
const b2 = `{"Id":2,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"10896","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"039785","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
const b3 = `{"Id":3,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"14532","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049352","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
const b4 = `{"Id":4,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"16515","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049745","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
const b5 = `{"Id":5,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"14776","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049746","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
const b6 = `{"Id":6,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"15711","Kurzbezeichnung":"UniCredit Bank-HypoVereinb","Pan":"k2201","BIC":"4HYVEDEMM48","Pruefzifferberechnungsmethode":"89","Datensatznummer":"904974","Aenderungskennzeichen":"7","Bankleitzahlloeschung":"U","Nachfolgebankleitzahl":"00000000"}`

func TestQueryByBlzMatchesOneBank(t *testing.T) {

	// When: blz is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?blz=10010424", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	// Then: status is ok
	AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := `{"Banks":[` + b1 + "]}"
	AssertEquals(t, expected, rr.Body.String())
}

func TestQueryByBlzMatchesMoreBanks(t *testing.T) {

	// When: blz is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?blz=10020890", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	// Then: status is ok
	AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 5 matching bank
	expected := `{"Banks":[` + b2 + "," + b3 + "," + b4 + "," + b5 + "," + b6 + "]}"
	AssertEquals(t, expected, rr.Body.String())
}

func TestNotFound(t *testing.T) {

	// When: blz unknown is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?blz=1002089", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	// Then: status is ok
	AssertEquals(t, http.StatusNotFound, rr.Code)

	// And: Respose body is empty
	expected := ""
	AssertEquals(t, expected, rr.Body.String())
}

func TestQueryByBicMatchesOneBank(t *testing.T) {

	// When: bic is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?bic=AARBDE5W100", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	// Then: status is ok
	AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := `{"Banks":[{"Id":1,"Blz":"10010424","Bankleitzahlfuehrend":"","Bezeichnung":"Aareal Bank","PLZ":"10666","Kurzbezeichnung":"Aareal Bank","Pan":"26910","BIC":"AARBDE5W100","Pruefzifferberechnungsmethode":"09","Datensatznummer":"004795","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}]}`
	AssertEquals(t, expected, rr.Body.String())
}

func TestQueryByNameMatchesOneBank(t *testing.T) {

	// When: name is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?name=Aareal+Bank", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	// Then: status is ok
	AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := `{"Banks":[{"Id":1,"Blz":"10010424","Bankleitzahlfuehrend":"","Bezeichnung":"Aareal Bank","PLZ":"10666","Kurzbezeichnung":"Aareal Bank","Pan":"26910","BIC":"AARBDE5W100","Pruefzifferberechnungsmethode":"09","Datensatznummer":"004795","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}]}`
	AssertEquals(t, expected, rr.Body.String())
}

func TestGetById(t *testing.T) {

	// When: bank with id 1 is gotten
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks/1", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	// Then: status is ok
	AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := b1
	AssertEquals(t, expected, rr.Body.String())
}

func TestSerializeBankResponse(t *testing.T) {
	response := ResponseWrapper{Banks: []bank.Bank{{Blz: "12345"}}}
	json := Json(response)
	expected := `{"Banks":[{"Id":0,"Blz":"12345","Bankleitzahlfuehrend":"","Bezeichnung":"","PLZ":"","Kurzbezeichnung":"","Pan":"","BIC":"","Pruefzifferberechnungsmethode":"","Datensatznummer":"","Aenderungskennzeichen":"","Bankleitzahlloeschung":"","Nachfolgebankleitzahl":""}]}`
	AssertEquals(t, expected, json)
}

func AssertEquals(t *testing.T, expect interface{}, actual interface{}) {
	if expect != actual {
		t.Errorf("wanted: %v, \ngot: %v", expect, actual)
		t.FailNow()
	}
}
