package parser

import (
	"testing"
	"bitbucket.org/rwirdemann/bundesbank/util"
)

const line = "100104241Aareal Bank                                               10666Berlin                             Aareal Bank                26910AARBDE5W10009004795U000000000"

func TestParse(t *testing.T) {
	bank := parseLine(line)
	util.AssertEquals(t, "10010424", bank.Blz)
	util.AssertEquals(t, "", bank.Bankleitzahlfuehrend)
	util.AssertEquals(t, "Aareal Bank", bank.Bezeichnung)
	util.AssertEquals(t, "10666", bank.PLZ)
	util.AssertEquals(t, "Aareal Bank", bank.Kurzbezeichnung)
	util.AssertEquals(t, "26910", bank.Pan)
	util.AssertEquals(t, "AARBDE5W100", bank.BIC)
	util.AssertEquals(t, "09", bank.Pruefzifferberechnungsmethode)
	util.AssertEquals(t, "004795", bank.Datensatznummer)
	util.AssertEquals(t, "U", bank.Aenderungskennzeichen)
	util.AssertEquals(t, "0", bank.Bankleitzahlloeschung)
	util.AssertEquals(t, "00000000", bank.Nachfolgebankleitzahl)
}
