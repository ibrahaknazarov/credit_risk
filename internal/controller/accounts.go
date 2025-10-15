package controller

import (
	"github.com/ibrahaknzarov/credit_risk/internal/models/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AccountResponse struct {
	ID            int       `json:"id"`
	FirtName      string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	AccountStatus string    `json:"account_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (ctrl *Controller) GetAccountByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	r, err := ctrl.service.GetAccountByID(id)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, AccountResponse{
		ID:            r.ID,
		FirtName:      r.FirtName,
		LastName:      r.LastName,
		Email:         r.Email,
		Password:      r.Password,
		AccountStatus: r.AccountStatus,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	})
}

func (ctrl *Controller) GetAllAccounts(c *gin.Context) {
	r, err := ctrl.service.GetAllAccounts()
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	var respAccount []AccountResponse
	for _, account := range r {
		respAccount = append(respAccount, AccountResponse{
			ID:            account.ID,
			FirtName:      account.FirtName,
			LastName:      account.LastName,
			Email:         account.Email,
			Password:      account.Password,
			AccountStatus: account.AccountStatus,
			CreatedAt:     account.CreatedAt,
			UpdatedAt:     account.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, AccountResponse)
}

type CreateAccountRequest struct {
	ID            int    `json:"id"`
	FirtName      string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	AccountStatus string `json:"account_status"`
}

func (ctrl *Controller) CreateAccount(c *gin.Context) {
	var input CreateAccountRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, err)
		return
	}

	if err := ctrl.service.CreateAccount(domain.Room{
		ID:            input.ID,
		FirtName:      input.FirtName,
		LastName:      input.LastName,
		Email:         input.Email,
		Password:      input.Password,
		AccountStatus: input.AccountStatus,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Account created successfully!"})

}

type UpdateAccountRequest struct {
	ID            int    `json:"id"`
	FirtName      string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	AccountStatus string `json:"account_status"`
}

func (ctrl *Controller) UpdateAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	var input UpdateAccountRequest
	if err = c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, err)
		return
	}

	ctrl.service.UpdateAccountByID(domain.Account{
		ID:            id,
		FirtName:      input.FirtName,
		LastName:      input.LastName,
		Email:         input.Email,
		Password:      input.Password,
		AccountStatus: input.AccountStatus,
		UpdatedAt:     time.Now(),
	})

	c.JSON(http.StatusOK, CommonResponse{Message: "Account updated successfully!"})
}

func (ctrl *Controller) DeleteAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	if err = ctrl.service.DeleteAccountByID(id); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Account deleted successfully!"})
}
