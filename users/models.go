package users

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/nyebarid/ny-rest-api/common"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique_index"`
	Provider string
	Avatar   string
	SocialID string
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&User{})
}

func FindOneUser(condition interface{}) (User, error) {
	db := common.GetDB()
	var model User
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func (model *User) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}
