package service
import (
	"credit_risk/internal/contracts"
	"credit_risk/internal/repository"
)
type Service struct {
	reppository contracts.Repository
}

func NewService(repo contracts.Repository) *Service {
	return &Service{reppository: repo}
}
