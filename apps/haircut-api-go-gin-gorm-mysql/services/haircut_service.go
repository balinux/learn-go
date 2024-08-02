package services

import (
	"haircut-api-go-gin-gorm-mysql/models"
	"haircut-api-go-gin-gorm-mysql/repositories"
)

type HaircutService interface {
	GetAllHaircuts() ([]models.Haircut, error)
	GetHaircutByID(id uint) (models.Haircut, error)
	CreateHaircut(haircut models.Haircut) (models.Haircut, error)
	UpdateHaircut(haircut models.Haircut) (models.Haircut, error)
	DeleteHaircut(id uint) error
}

type haircutService struct {
	repository repositories.HaircutRepository
}

// CreateHaircut implements HaircutService.
func (h *haircutService) CreateHaircut(haircut models.Haircut) (models.Haircut, error) {
	panic("unimplemented")
}

// DeleteHaircut implements HaircutService.
func (h *haircutService) DeleteHaircut(id uint) error {
	panic("unimplemented")
}

// GetAllHaircuts implements HaircutService.
func (h *haircutService) GetAllHaircuts() ([]models.Haircut, error) {
	return h.repository.FindAll()
}

// GetHaircutByID implements HaircutService.
func (h *haircutService) GetHaircutByID(id uint) (models.Haircut, error) {
	panic("unimplemented")
}

// UpdateHaircut implements HaircutService.
func (h *haircutService) UpdateHaircut(haircut models.Haircut) (models.Haircut, error) {
	panic("unimplemented")
}

func NewHaircutService(repository repositories.HaircutRepository) HaircutService {
	return &haircutService{repository}
}
