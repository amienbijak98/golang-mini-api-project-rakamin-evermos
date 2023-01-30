package repositories

import (
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(db *gorm.DB, user models.User) error {
	hash, _ := HashPassword(user.Password)
	user.Password = hash

	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	store := &models.Store{
		UserID:    user.ID,
		StoreName: "My Store",
		PhotoUrl:  "",
	}

	result = db.Create(&store)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CheckEmail() {

}

func CheckMobilePhone() {

}
