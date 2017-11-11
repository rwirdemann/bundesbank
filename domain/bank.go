package domain

type Bank struct {
	Blz                           string
	Bankleitzahlfuehrend          string
	Bezeichnung                   string
	PLZ                           string
	Kurzbezeichnung               string
	Pan                           string
	BIC                           string
	Pruefzifferberechnungsmethode string
	Datensatznummer               string
	Aenderungskennzeichen         string
	Bankleitzahlloeschung         string
	Nachfolgebankleitzahl         string
}
