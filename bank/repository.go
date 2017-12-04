package bank

type Repository interface {
	NextId() int
	Add(bank Bank)
	ByBlz(blz string) ([]Bank, bool)
	ByBic(bic string) ([]Bank, bool)
	ByBezeichnung(bezeichnung string) ([]Bank, bool)
	ById(id int) (Bank, bool)
}

type MapBasedRepository struct {
	id                 int
	banksById          map[int]Bank
	banksByBlz         map[string][]Bank
	banksByBic         map[string][]Bank
	banksByBezeichnung map[string][]Bank
}

type FileRepository struct {
	MapBasedRepository
}

func NewFileRepository() *FileRepository {
	r := FileRepository{}
	r.banksById = make(map[int]Bank)
	r.banksByBezeichnung = make(map[string][]Bank)
	r.banksByBic = make(map[string][]Bank)
	r.banksByBlz = make(map[string][]Bank)
	return &r
}

func (c *MapBasedRepository) ById(id int) (Bank, bool) {
	bank, ok := c.banksById[id]
	return bank, ok
}

func (c *MapBasedRepository) ByBlz(blz string) ([]Bank, bool) {
	banks, ok := c.banksByBlz[blz]
	return banks, ok
}

func (c *MapBasedRepository) ByBic(bic string) ([]Bank, bool) {
	banks, ok := c.banksByBic[bic]
	return banks, ok
}

func (c *MapBasedRepository) ByBezeichnung(bezeichnung string) ([]Bank, bool) {
	banks, ok := c.banksByBezeichnung[bezeichnung]
	return banks, ok
}

func (c *MapBasedRepository) NextId() int {
	c.id++
	return c.id
}

func (c *MapBasedRepository) Add(bank Bank) {
	c.banksById[bank.Id] = bank
	c.addBankToBezeichnungMap(bank)
	c.addBankToBicMap(bank)
	c.addBankToPlzMap(bank)
}

func (c *MapBasedRepository) addBankToBezeichnungMap(bank Bank) {
	if bankArray, ok := c.banksByBezeichnung[bank.Bezeichnung]; ok {
		c.banksByBezeichnung[bank.Bezeichnung] = append(bankArray, bank)
	} else {
		bankArray := []Bank{bank}
		c.banksByBezeichnung[bank.Bezeichnung] = bankArray
	}
}

func (c *MapBasedRepository) addBankToBicMap(bank Bank) {
	if bank.BIC != "" {
		if bankArray, ok := c.banksByBic[bank.BIC]; ok {
			c.banksByBic[bank.BIC] = append(bankArray, bank)
		} else {
			bankArray := []Bank{bank}
			c.banksByBic[bank.BIC] = bankArray
		}
	}
}

func (c *MapBasedRepository) addBankToPlzMap(bank Bank) {
	if bankArray, ok := c.banksByBlz[bank.Blz]; ok {
		c.banksByBlz[bank.Blz] = append(bankArray, bank)
	} else {
		c.banksByBlz[bank.Blz] = []Bank{bank}
	}
}
