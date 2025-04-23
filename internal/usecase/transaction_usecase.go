// Placeholder for transaction_usecase.go
package usecase

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "sync"
    "kreditplus/internal/domain"
    "kreditplus/internal/repository"
)

type TransactionUsecase struct {
    Repo *repository.TransactionRepository
    Mu   sync.Mutex
}

func NewTransactionUsecase(repo *repository.TransactionRepository) *TransactionUsecase {
    return &TransactionUsecase{
        Repo: repo,
    }
}

func (u *TransactionUsecase) CreateTransaction(c *gin.Context) {
    u.Mu.Lock()
    defer u.Mu.Unlock()

    var trx domain.Transaction
    if err := c.ShouldBindJSON(&trx); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := u.Repo.Create(&trx); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, trx)
}
