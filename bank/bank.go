package bank

type Bank struct {
	Id                            int
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
