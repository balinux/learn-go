package repositories

import (
	"haircut-api-go-gin-gorm-mysql/models"

	"gorm.io/gorm"
)

type HaircutRepository interface {
	FindAll() ([]models.HairCut, error)
	FindByID(id uint) (models.HairCut, error)
	Create(haircut models.HairCut) (models.HairCut, error)
	Update(haircut models.HairCut) (models.HairCut, error)
	Delete(haircut models.HairCut) error
}

type haircutRepository struct {
	db *gorm.DB
}

func NewHaircutRepository(db *gorm.DB) HaircutRepository {
	return &haircutRepository{db}
}

// Create implements HaircutRepository.
func (h *haircutRepository) Create(haircut models.HairCut) (models.HairCut, error) {
	panic("unimplemented")
}

// Delete implements HaircutRepository.
func (h *haircutRepository) Delete(haircut models.HairCut) error {
	panic("unimplemented")
}

// FindAll implements HaircutRepository.
func (h *haircutRepository) FindAll() ([]models.HairCut, error) {
	var haircuts []models.HairCut
	err := h.db.Find(&haircuts).Error
	return haircuts, err
}

// FindByID implements HaircutRepository.
func (h *haircutRepository) FindByID(id uint) (models.HairCut, error) {
	panic("unimplemented")
}

// Update implements HaircutRepository.
func (h *haircutRepository) Update(haircut models.HairCut) (models.HairCut, error) {
	panic("unimplemented")
}
