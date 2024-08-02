package repositories

import (
	"haircut-api-go-gin-gorm-mysql/models"

	"gorm.io/gorm"
)

type HaircutRepository interface {
	FindAll() ([]models.Haircut, error)
	FindByID(id uint) (models.Haircut, error)
	Create(haircut models.Haircut) (models.Haircut, error)
	Update(haircut models.Haircut) (models.Haircut, error)
	Delete(haircut models.Haircut) error
}

type haircutRepository struct {
	db *gorm.DB
}

func NewHaircutRepository(db *gorm.DB) HaircutRepository {
	return &haircutRepository{db}
}

// Create implements HaircutRepository.
func (h *haircutRepository) Create(haircut models.Haircut) (models.Haircut, error) {
	panic("unimplemented")
}

// Delete implements HaircutRepository.
func (h *haircutRepository) Delete(haircut models.Haircut) error {
	panic("unimplemented")
}

// FindAll implements HaircutRepository.
func (h *haircutRepository) FindAll() ([]models.Haircut, error) {
	var haircuts []models.Haircut
	err := h.db.Find(&haircuts).Error
	return haircuts, err
}

// FindByID implements HaircutRepository.
func (h *haircutRepository) FindByID(id uint) (models.Haircut, error) {
	panic("unimplemented")
}

// Update implements HaircutRepository.
func (h *haircutRepository) Update(haircut models.Haircut) (models.Haircut, error) {
	panic("unimplemented")
}
