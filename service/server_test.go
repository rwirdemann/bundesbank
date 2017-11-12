package service

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bitbucket.org/rwirdemann/bundesbank/util"
	"bitbucket.org/rwirdemann/bundesbank/domain"
)

func init()  {
	Repository = domain.GetRepositoryInstance()
	ImportBundesbankFile("service_test_data.txt")
}

func TestQueryByBlzMatchesOneBank(t *testing.T) {

	// When: blz is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?blz=10010424", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(banks)
	handler.ServeHTTP(rr, req)

	// Then: status is ok
	util.AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := `{"Banks":[{"Id":1,"Blz":"10010424","Bankleitzahlfuehrend":"","Bezeichnung":"Aareal Bank","PLZ":"10666","Kurzbezeichnung":"Aareal Bank","Pan":"26910","BIC":"AARBDE5W100","Pruefzifferberechnungsmethode":"09","Datensatznummer":"004795","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}]}`
	util.AssertEquals(t, expected, rr.Body.String())
}

func TestQueryByBlzMatchesMoreBanks(t *testing.T) {

	// When: blz is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?blz=10020890", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(banks)
	handler.ServeHTTP(rr, req)

	// Then: status is ok
	util.AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := `{"Banks":[{"Id":2,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"10896","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"039785","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"},{"Id":3,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"14532","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049352","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"},{"Id":4,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"16515","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049745","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"},{"Id":5,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"14776","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049746","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"},{"Id":6,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"15711","Kurzbezeichnung":"UniCredit Bank-HypoVereinb","Pan":"k2201","BIC":"4HYVEDEMM48","Pruefzifferberechnungsmethode":"89","Datensatznummer":"904974","Aenderungskennzeichen":"7","Bankleitzahlloeschung":"U","Nachfolgebankleitzahl":"00000000"}]}`
	util.AssertEquals(t, expected, rr.Body.String())
}

func TestNotFound(t *testing.T) {

	// When: blz unknown is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?blz=1002089", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(banks)
	handler.ServeHTTP(rr, req)

	// Then: status is ok
	util.AssertEquals(t, http.StatusNotFound, rr.Code)

	// And: Respose body is empty
	expected := ""
	util.AssertEquals(t, expected, rr.Body.String())
}

func TestQueryByBicMatchesOneBank(t *testing.T) {

	// When: bic is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?bic=AARBDE5W100", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(banks)
	handler.ServeHTTP(rr, req)

	// Then: status is ok
	util.AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := `{"Banks":[{"Id":1,"Blz":"10010424","Bankleitzahlfuehrend":"","Bezeichnung":"Aareal Bank","PLZ":"10666","Kurzbezeichnung":"Aareal Bank","Pan":"26910","BIC":"AARBDE5W100","Pruefzifferberechnungsmethode":"09","Datensatznummer":"004795","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}]}`
	util.AssertEquals(t, expected, rr.Body.String())
}

func TestQueryByNameMatchesOneBank(t *testing.T) {

	// When: name is queried
	req, _ := http.NewRequest("GET", "/bundesbank/v1/banks?name=Aareal+Bank", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(banks)
	handler.ServeHTTP(rr, req)

	// Then: status is ok
	util.AssertEquals(t, http.StatusOK, rr.Code)

	// And: Body contains 1 matching bank
	expected := `{"Banks":[{"Id":1,"Blz":"10010424","Bankleitzahlfuehrend":"","Bezeichnung":"Aareal Bank","PLZ":"10666","Kurzbezeichnung":"Aareal Bank","Pan":"26910","BIC":"AARBDE5W100","Pruefzifferberechnungsmethode":"09","Datensatznummer":"004795","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}]}`
	util.AssertEquals(t, expected, rr.Body.String())
}