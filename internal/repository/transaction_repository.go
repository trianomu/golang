package repository

import (
    "gorm.io/gorm"
    "kreditplus/internal/domain"
)

type TransactionRepository struct {
    DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
    return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) Create(tx *domain.Transaction) error {
    return r.DB.Create(tx).Error
}
