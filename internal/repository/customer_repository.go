package repository

import (
    "gorm.io/gorm"
    "kreditplus/internal/domain"
)

type CustomerRepository struct {
    DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
    return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) Create(customer *domain.Customer) error {
    return r.DB.Create(customer).Error
}

func (r *CustomerRepository) FindByID(id string) (*domain.Customer, error) {
    var customer domain.Customer
    err := r.DB.First(&customer, "id = ?", id).Error
    return &customer, err
}
