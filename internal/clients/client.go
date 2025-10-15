package clients

import (
	"github.com/ibrahaknazarov/credit_risk/internal/models/domain"
)

type Client struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func NewClient(id int, firstName, lastName, email, phone string) *Client {
	return &Client{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
}

