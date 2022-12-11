package payement


import (
	"gorm.io/gorm"
)

type IPayementRepository interface {
	FindAll() ([]Payement, error)
	FindById(id int) (Payement, error)
	Create(payement Payement) (Payement, error)
	Update(payement Payement) (Payement, error)
	Delete(id int) error
	Stream() (Payement, error)
}

type PayementRepository struct {
	db *gorm.DB
}

func NewPayementRepository(db *gorm.DB) *PayementRepository {
	return &PayementRepository{db}
}

