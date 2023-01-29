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
	db.First(&category, id)
	return category, nil
}

func UpdateCategory(db *gorm.DB, id int, category models.Category) (models.Category, error) {
	var existingCategory models.Category
	db.First(&existingCategory, id)

	existingCategory.CategoryName = category.CategoryName
	existingCategory.UpdatedAt = time.Now()

	db.Save(&existingCategory)
	return existingCategory, nil
}

func DeleteCategory(db *gorm.DB, id int) error {
	var category models.Category
	db.First(&category, id)

	db.Delete(&category)
	return nil
}
