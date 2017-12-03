package bank

import (
	"testing"
)

const line = "100104241Aareal Bank                                               10666Berlin                             Aareal Bank                26910AARBDE5W10009004795U000000000"

func TestParse(t *testing.T) {
	bank := parseLine(line)
	AssertEquals(t, "10010424", bank.Blz)
	AssertEquals(t, "", bank.Bankleitzahlfuehrend)
	AssertEquals(t, "Aareal Bank", bank.Bezeichnung)
	AssertEquals(t, "10666", bank.PLZ)
	AssertEquals(t, "Aareal Bank", bank.Kurzbezeichnung)
	AssertEquals(t, "26910", bank.Pan)
	AssertEquals(t, "AARBDE5W100", bank.BIC)
	AssertEquals(t, "09", bank.Pruefzifferberechnungsmethode)
	AssertEquals(t, "004795", bank.Datensatznummer)
	AssertEquals(t, "U", bank.Aenderungskennzeichen)
	AssertEquals(t, "0", bank.Bankleitzahlloeschung)
	AssertEquals(t, "00000000", bank.Nachfolgebankleitzahl)
}

func AssertEquals(t *testing.T, expect interface{}, actual interface{}) {
	if expect != actual {
		t.Errorf("wanted: %v, \ngot: %v", expect, actual)
		t.FailNow()
	}
}

