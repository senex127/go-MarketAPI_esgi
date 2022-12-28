package payement


import (
	"gorm.io/gorm"
)

type Repository interface {
	// FindAll() ([]Payement, error)
	// FindById(id int) (Payement, error)
	Create(payement Payement) (Payement, error)
	// Update(payement Payement) (Payement, error)
	// Delete(id int) error
	// Stream() (Payement, error)
}

type repository struct {
	db *gorm.DB
}

func NewPayementRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) Create(payement Payement) (Payement, error) {
	repo.db.Create(&payement)

	// save payement in db
	repo.db.Save(&payement)
	return payement, nil
}

