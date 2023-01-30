package repositories

import (
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// birthDate, _ := time.Parse("2006-01-02", user.BirthDate.String())
	// user.BirthDate = birthDate

	result := db.Clauses(clause.Returning{}).Create(&user)
	if result.Error != nil {
		return result.Error
	}

	store := &models.Store{
		UserID:    user.ID,
		StoreName: user.Name + "'s Store",
		PhotoUrl:  "",
	}

	result = db.Create(&store)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func CheckEmail(db *gorm.DB, user models.User) error {
// 	result := db.Where("email = ?", user.Password).First(user)
// 	if result.Error == nil {
// 		return errors.New("")
// 	}
// 	return nil
// }

// func CheckMobilePhone() {

// }
