package domain

import "sync"

type BankRepository interface {
	NextId() int
	Add(bank Bank)
	ByBlz(blz string)([]Bank, bool)
	ByBic(bic string)([]Bank, bool)
	ByBezeichnung(bezeichnung string)([]Bank, bool)
}

type BankRepositoryMemory struct {
	id                 int
	banksByBlz         map[string][]Bank
	banksByBic         map[string][]Bank
	banksByBezeichnung map[string][]Bank
}

func (c *BankRepositoryMemory) ByBlz(blz string) ([]Bank, bool) {
	banks, ok := c.banksByBlz[blz]
	return banks, ok
}

func (c *BankRepositoryMemory) ByBic(bic string) ([]Bank, bool) {
	banks, ok := c.banksByBic[bic]
	return banks, ok
}

func (c *BankRepositoryMemory) ByBezeichnung(bezeichnung string) ([]Bank, bool) {
	banks, ok := c.banksByBezeichnung[bezeichnung]
	return banks, ok
}

func (c *BankRepositoryMemory) NextId() int {
	c.id++
	return c.id
}

func (c *BankRepositoryMemory) Add(bank Bank) {
	c.addBankToBezeichnungMap(bank)
	c.addBankToBicMap(bank)
	c.addBankToPlzMap(bank)
}

func (c *BankRepositoryMemory) addBankToBezeichnungMap(bank Bank) {
	if bankArray, ok := c.banksByBezeichnung[bank.Bezeichnung]; ok {
		c.banksByBezeichnung[bank.Bezeichnung] = append(bankArray, bank)
	} else {
		bankArray := []Bank{bank}
		c.banksByBezeichnung[bank.Bezeichnung] = bankArray
	}
}

func (c *BankRepositoryMemory) addBankToBicMap(bank Bank) {
	if bank.BIC != "" {
		if bankArray, ok := c.banksByBic[bank.BIC]; ok {
			c.banksByBic[bank.BIC] = append(bankArray, bank)
		} else {
			bankArray := []Bank{bank}
			c.banksByBic[bank.BIC] = bankArray
		}
	}
}

func (c *BankRepositoryMemory) addBankToPlzMap(bank Bank) {
	if bankArray, ok := c.banksByBlz[bank.Blz]; ok {
		c.banksByBlz[bank.Blz] = append(bankArray, bank)
	} else {
		c.banksByBlz[bank.Blz] = []Bank{bank}
	}
}

var repository BankRepository
var once sync.Once

func GetRepositoryInstance() BankRepository {
	once.Do(func() {
		r := &BankRepositoryMemory{}
		r.banksByBezeichnung = make(map[string][]Bank)
		r.banksByBic = make(map[string][]Bank)
		r.banksByBlz = make(map[string][]Bank)
		repository = r
	})
	return repository
}
