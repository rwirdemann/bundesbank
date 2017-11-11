package service

import (
	"strings"
	"os"
	"bufio"
	"bitbucket.org/rwirdemann/bundesbank/domain"
)

type Field struct {
	Start int
	End   int
}

var BanksByPlz map[string][]domain.Bank
var BanksByBic map[string][]domain.Bank
var BanksByBezeichnung map[string][]domain.Bank

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
	BanksByPlz = make(map[string][]domain.Bank)
	BanksByBic = make(map[string][]domain.Bank)
	BanksByBezeichnung = make(map[string][]domain.Bank)

	if file, err := os.Open(file); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			bank := parseLine(scanner.Text())
			addBankToPlzMap(bank)
			addBankToBicMap(bank)
			addBankToBezeichnungMap(bank)
		}
	} else {
		panic(err)
	}
}

func addBankToBezeichnungMap(bank domain.Bank) {
	if bankArray, ok := BanksByBezeichnung[bank.Bezeichnung]; ok {
		BanksByBezeichnung[bank.Bezeichnung] = append(bankArray, bank)
	} else {
		bankArray := []domain.Bank{bank}
		BanksByBezeichnung[bank.Bezeichnung] = bankArray
	}
}

func addBankToBicMap(bank domain.Bank) {
	if bank.BIC != "" {
		if bankArray, ok := BanksByBic[bank.BIC]; ok {
			BanksByBic[bank.BIC] = append(bankArray, bank)
		} else {
			bankArray := []domain.Bank{bank}
			BanksByBic[bank.BIC] = bankArray
		}
	}
}

func addBankToPlzMap(bank domain.Bank) {
	if bankArray, ok := BanksByPlz[bank.Blz]; ok {
		BanksByPlz[bank.Blz] = append(bankArray, bank)
	} else {
		BanksByPlz[bank.Blz] = []domain.Bank{bank}
	}
}

func parseLine(line string) domain.Bank {
	var b domain.Bank
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
