package controller

import (
	"errors"
	"github.com/ibrahaknazarov/credit_risk/internal/contracts"
	"github.com/ibrahaknazarov/credit_risk/internal/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service contracts.ServiceI
}

func NewController(svc contracts.ServiceI) *Controller {
	return &Controller{
		service: svc,
	}
}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrProductNotfound) ||
		errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidProductID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword) || errors.Is(err, errs.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue) ||
		errors.Is(err, errs.ErrInvalidProductName) ||
		errors.Is(err, errs.ErrUsernameAlreadyExists):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
