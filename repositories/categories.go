package repositories

import (
	"time"

	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) error {
	result := db.Create(&category)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllCategories(db *gorm.DB) ([]models.Category, error) {
	var categories []models.Category

	result := db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func GetCategoryByID(db *gorm.DB, id int) (models.Category, error) {
	var category models.Category
	result := db.First(&category, id)
	if result.Error != nil {
		return models.Category{}, result.Error
	}
	return category, nil
}

func UpdateCategory(db *gorm.DB, id int, category models.Category) (models.Category, error) {
	var existingCategory models.Category
	result := db.First(&existingCategory, id)
	if result.Error != nil {
		return models.Category{}, result.Error
	}

	existingCategory.CategoryName = category.CategoryName
	existingCategory.UpdatedAt = time.Now()

	result = db.Save(&existingCategory)
	if result.Error != nil {
		return models.Category{}, result.Error
	}
	return existingCategory, nil
}

func DeleteCategory(db *gorm.DB, id int) error {
	var category models.Category
	result := db.First(&category, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
