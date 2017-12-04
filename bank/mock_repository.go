package bank

import (
	"encoding/json"
)

type MockRepository struct {
	id                 int
	banksById          map[int]Bank
	banksByBlz         map[string][]Bank
	banksByBic         map[string][]Bank
	banksByBezeichnung map[string][]Bank
}

func NewMockRepository() Repository {
	r := MockRepository{}
	r.banksById = make(map[int]Bank)
	r.banksByBezeichnung = make(map[string][]Bank)
	r.banksByBic = make(map[string][]Bank)
	r.banksByBlz = make(map[string][]Bank)

	const b1 = `{"Id":1,"Blz":"10010424","Bankleitzahlfuehrend":"","Bezeichnung":"Aareal Bank","PLZ":"10666","Kurzbezeichnung":"Aareal Bank","Pan":"26910","BIC":"AARBDE5W100","Pruefzifferberechnungsmethode":"09","Datensatznummer":"004795","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
	const b2 = `{"Id":2,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"10896","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"039785","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
	const b3 = `{"Id":3,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"14532","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049352","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
	const b4 = `{"Id":4,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"16515","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049745","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
	const b5 = `{"Id":5,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"14776","Kurzbezeichnung":"UniCredit Bank-HypoVereinbk","Pan":"22014","BIC":"HYVEDEMM488","Pruefzifferberechnungsmethode":"99","Datensatznummer":"049746","Aenderungskennzeichen":"U","Bankleitzahlloeschung":"0","Nachfolgebankleitzahl":"00000000"}`
	const b6 = `{"Id":6,"Blz":"10020890","Bankleitzahlfuehrend":"","Bezeichnung":"UniCredit Bank - HypoVereinsbank","PLZ":"15711","Kurzbezeichnung":"UniCredit Bank-HypoVereinb","Pan":"k2201","BIC":"4HYVEDEMM48","Pruefzifferberechnungsmethode":"89","Datensatznummer":"904974","Aenderungskennzeichen":"7","Bankleitzahlloeschung":"U","Nachfolgebankleitzahl":"00000000"}`

	r.Add(unmarshal(b1))
	r.Add(unmarshal(b2))
	r.Add(unmarshal(b3))
	r.Add(unmarshal(b4))
	r.Add(unmarshal(b5))
	r.Add(unmarshal(b6))

	return &r
}

func unmarshal(s string) Bank {
	b := Bank{}
	json.Unmarshal([]byte(s), &b)
	return b
}

func (c *MockRepository) ById(id int) (Bank, bool) {
	bank, ok := c.banksById[id]
	return bank, ok
}

func (c *MockRepository) ByBlz(blz string) ([]Bank, bool) {
	banks, ok := c.banksByBlz[blz]
	return banks, ok
}

func (c *MockRepository) ByBic(bic string) ([]Bank, bool) {
	banks, ok := c.banksByBic[bic]
	return banks, ok
}

func (c *MockRepository) ByBezeichnung(bezeichnung string) ([]Bank, bool) {
	banks, ok := c.banksByBezeichnung[bezeichnung]
	return banks, ok
}

func (c *MockRepository) NextId() int {
	c.id++
	return c.id
}

func (c *MockRepository) Add(bank Bank) {
	c.banksById[bank.Id] = bank
	c.addBankToBezeichnungMap(bank)
	c.addBankToBicMap(bank)
	c.addBankToPlzMap(bank)
}

func (c *MockRepository) addBankToBezeichnungMap(bank Bank) {
	if bankArray, ok := c.banksByBezeichnung[bank.Bezeichnung]; ok {
		c.banksByBezeichnung[bank.Bezeichnung] = append(bankArray, bank)
	} else {
		bankArray := []Bank{bank}
		c.banksByBezeichnung[bank.Bezeichnung] = bankArray
	}
}

func (c *MockRepository) addBankToBicMap(bank Bank) {
	if bank.BIC != "" {
		if bankArray, ok := c.banksByBic[bank.BIC]; ok {
			c.banksByBic[bank.BIC] = append(bankArray, bank)
		} else {
			bankArray := []Bank{bank}
			c.banksByBic[bank.BIC] = bankArray
		}
	}
}

func (c *MockRepository) addBankToPlzMap(bank Bank) {
	if bankArray, ok := c.banksByBlz[bank.Blz]; ok {
		c.banksByBlz[bank.Blz] = append(bankArray, bank)
	} else {
		c.banksByBlz[bank.Blz] = []Bank{bank}
	}
}
