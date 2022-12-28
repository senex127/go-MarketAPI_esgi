package payement

import "time"

type Service interface {
	// FindAll() ([]Payement, error)
	// FindById(id int) (Payement, error)
	Create(productId int, price string) (Payement, error)
	// Update(payement Payement) (Payement, error)
	// Delete(id int) error
	// Stream() (Payement, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Create(productId int, price string ) (Payement, error) {

	//WORK IN PROGRESS but insert in db works
	payementObject := Payement{
		ProductID: productId,
		PricePaid: price,
		CreatedAt: time.Now().Format("2006-01-02"),
		UpdatedAt: time.Now().Format("2006-01-02"),
	}
	
	s.repo.Create(payementObject)
	
	
	return payementObject, nil
}

// func (s *service) FindAll() ([]Payement, error) {
// 	//return s.repo.FindAll()
// }


// func (s *service) FindById(id int) (Payement, error) {
// 	//return s.repo.FindById(id)
// }

// func (s *service) Create(payement Payement) (Payement, error) {
// 	//return s.repo.Create(payement)
// }


// func (s *service) Update(payement Payement) (Payement, error) {
// 	//return s.repo.Update(payement)
// }

// func (s *service) Delete(id int) error {
// 	//return s.repo.Delete(id)
// }

// func (s *service) Stream() (Payement, error) {
// 	//return s.repo.Stream()
// }