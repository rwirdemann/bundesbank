package bank

import (
	"strings"
	"os"
	"bufio"
)

type Field struct {
	Start int
	End   int
}

var Fields = map[string]Field{
	"blz":                           {Start: 0, End: 8},
	"blzfuehrend":                   {Start: 8, End: 9},
	"bezeichnung":                   {Start: 9, End: 67},
	"plz":                           {Start: 67, End: 72},
	"kurzbezeichnung":               {Start: 107, End: 134},
	"pan":                           {Start: 134, End: 139},
	"bic":                           {Start: 139, End: 150},
	"pruefzifferberechnungsmethode": {Start: 150, End: 152},
	"datensatznummer":               {Start: 152, End: 158},
	"aenderungskennzeichen":         {Start: 158, End: 159},
	"bankleitzahlloeschung":         {Start: 159, End: 160},
	"nachfolgebankleitzahl":         {Start: 160, End: 168},
}

func ImportBundesbankFile(file string, s *Service) {
	if file, err := os.Open(file); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			b := parseLine(scanner.Text())
			b.Id = s.BankRepository.NextId()
			s.BankRepository.Add(b)
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
