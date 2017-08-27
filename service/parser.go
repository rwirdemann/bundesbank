package service

import (
	"strings"
	"os"
	"bufio"
)

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

type Field struct {
	Start int
	End   int
}

var BanksByPlz map[string][]Bank
var BanksByBic map[string][]Bank
var BanksByBezeichnung map[string][]Bank

var Fields = map[string]Field{
	"blz":                           Field{Start: 0, End: 8},
	"blzfuehrend":                   Field{Start: 8, End: 9},
	"bezeichnung":                   Field{Start: 9, End: 67},
	"plz":                           Field{Start: 67, End: 72},
	"kurzbezeichnung":               Field{Start: 107, End: 134},
	"pan":                           Field{Start: 134, End: 139},
	"bic":                           Field{Start: 139, End: 150},
	"pruefzifferberechnungsmethode": Field{Start: 150, End: 152},
	"datensatznummer":               Field{Start: 152, End: 158},
	"aenderungskennzeichen":         Field{Start: 158, End: 159},
	"bankleitzahlloeschung":         Field{Start: 159, End: 160},
	"nachfolgebankleitzahl":         Field{Start: 160, End: 168},
}

func ImportBundesbankFile(file string) {
	BanksByPlz = make(map[string][]Bank)
	BanksByBic = make(map[string][]Bank)
	BanksByBezeichnung = make(map[string][]Bank)

	if file, err := os.Open(file); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			bank := parseLine(scanner.Text())
			if bankArray, ok := BanksByPlz[bank.Blz]; ok {
				BanksByPlz[bank.Blz] = append(bankArray, bank)
			} else {
				bankArray := []Bank{bank}
				BanksByPlz[bank.Blz] = bankArray
			}

			if bank.BIC != "" {
				if bankArray, ok := BanksByBic[bank.BIC]; ok {
					BanksByBic[bank.BIC] = append(bankArray, bank)
				} else {
					bankArray := []Bank{bank}
					BanksByBic[bank.BIC] = bankArray
				}
			}

			if bankArray, ok := BanksByBezeichnung[bank.Bezeichnung]; ok {
				BanksByBezeichnung[bank.Bezeichnung] = append(bankArray, bank)
			} else {
				bankArray := []Bank{bank}
				BanksByBezeichnung[bank.Bezeichnung] = bankArray
			}
		}
	} else {
		panic(err)
	}
}

func parseLine(line string) Bank {
	var b Bank
	b.Blz = line[Fields["blz"].Start:Fields["blz"].End]
	b.Bezeichnung = strings.Trim(line[Fields["bezeichnung"].Start:Fields["bezeichnung"].End], " ")
	b.PLZ = strings.Trim(line[Fields["plz"].Start:Fields["plz"].End], " ")
	b.Kurzbezeichnung = strings.Trim(line[Fields["kurzbezeichnung"].Start:Fields["kurzbezeichnung"].End], " ")
	b.Pan = strings.Trim(line[Fields["pan"].Start:Fields["pan"].End], " ")
	b.BIC = strings.Trim(line[Fields["bic"].Start:Fields["bic"].End], " ")
	b.Pruefzifferberechnungsmethode = strings.Trim(line[Fields["pruefzifferberechnungsmethode"].Start:Fields["pruefzifferberechnungsmethode"].End], " ")
	b.Datensatznummer = strings.Trim(line[Fields["datensatznummer"].Start:Fields["datensatznummer"].End], " ")
	b.Aenderungskennzeichen = strings.Trim(line[Fields["aenderungskennzeichen"].Start:Fields["aenderungskennzeichen"].End], " ")
	b.Bankleitzahlloeschung = strings.Trim(line[Fields["bankleitzahlloeschung"].Start:Fields["bankleitzahlloeschung"].End], " ")
	b.Nachfolgebankleitzahl = strings.Trim(line[Fields["nachfolgebankleitzahl"].Start:Fields["nachfolgebankleitzahl"].End], " ")

	return b
}
