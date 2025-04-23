// Placeholder for customer_usecase.go
package usecase

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "kreditplus/internal/repository"
    "kreditplus/internal/domain"
)

type CustomerUsecase struct {
    Repo *repository.CustomerRepository
}

func NewCustomerUsecase(repo *repository.CustomerRepository) *CustomerUsecase {
    return &CustomerUsecase{Repo: repo}
}

func (u *CustomerUsecase) CreateCustomer(c *gin.Context) {
    var customer domain.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := u.Repo.Create(&customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, customer)
}

func (u *CustomerUsecase) GetCustomerByID(c *gin.Context) {
    id := c.Param("id")
    customer, err := u.Repo.FindByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, customer)
}
