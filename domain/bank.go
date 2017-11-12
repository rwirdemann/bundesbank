package domain

import "sync"

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

type BankRepository interface {
	NextId() int
}

type BankRepositoryMemory struct {
	id int
}

func (c *BankRepositoryMemory) NextId() int {
	c.id++
	return c.id
}

var repository BankRepository
var once sync.Once

func GetRepositoryInstance() BankRepository {
	once.Do(func() {
		repository = &BankRepositoryMemory{}
	})
	return repository
}
