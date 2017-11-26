package bank

type Service struct {
	BankRepository *FileRepository
}

func NewBankService(r *FileRepository) *Service {
	return &Service{BankRepository: r}
}
