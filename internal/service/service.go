package service
import (
	"github.com/ibrahaknzarov/credit_risk/internal/contracts"
	"github.com/ibrahaknzarov/credit_risk/internal/repository"
)
type Service struct {
	reppository contracts.Repository
}

func NewService(repo contracts.Repository) *Service {
	return &Service{reppository: repo}
}