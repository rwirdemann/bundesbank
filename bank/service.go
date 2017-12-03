package bank

type Service struct {
	BankRepository *FileRepository
}

func NewBankService(r *FileRepository) *Service {
	return &Service{BankRepository: r}
}

func (s Service) byId(id int) (Bank, bool){
	return s.BankRepository.ById(id)
}
