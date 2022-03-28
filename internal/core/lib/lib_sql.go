package lib

import (
	"github.com/GolfPhumrapee/finance/internal/models"
	"gorm.io/gorm"
)

// Repository repository interface
type SQL interface {
	Create(database *gorm.DB, i interface{}) error
	Update(database *gorm.DB, i interface{}) error
	Delete(database *gorm.DB, i interface{}) error
}

type sql struct{}

// NewSQL new repository
func NewSQL() SQL {
	return &sql{}
}

// Create create record database
func (repo *sql) Create(database *gorm.DB, i interface{}) error {
	if m, ok := i.(models.ModelInterface); ok {
		m.Stamp()
	}

	if err := database.Create(i).Error; err != nil {
		return err
	}

	return nil
}

// Update update record database
func (repo *sql) Update(database *gorm.DB, i interface{}) error {
	if m, ok := i.(models.ModelInterface); ok {
		m.UpdateStamp()
	}

	if err := database.Save(i).Error; err != nil {
		return err
	}

	return nil
}

// Delete delete record database
func (repo *sql) Delete(database *gorm.DB, i interface{}) error {
	if m, ok := i.(models.ModelInterface); ok {
		m.DeleteStamp()
	}

	if err := database.Delete(i).Error; err != nil {
		return err
	}

	return nil
}
