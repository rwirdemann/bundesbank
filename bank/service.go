package bank

type Service struct {
	Repository Repository
}

func NewBankService(r Repository) *Service {
	return &Service{Repository: r}
}

func (s Service) byId(id int) (Bank, bool) {
	return s.Repository.ById(id)
}
