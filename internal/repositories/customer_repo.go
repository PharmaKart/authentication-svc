package repositories

import (
	"github.com/PharmaKart/authentication-svc/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) (uuid.UUID, error)
	GetCustomerByUserID(userID string) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) (uuid.UUID, error) {
	if err := r.db.Create(customer).Error; err != nil {
		return uuid.Nil, err
	}
	return customer.ID, nil
}

func (r *customerRepository) GetCustomerByUserID(userID string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("user_id = ?", userID).First(&customer).Error
	return &customer, err
}
