package users

import (
	"time"

	"gitlab.com/nyebarid/ny-rest-api/common"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Role      bool       `gorm:"default:0" json:"role"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email"`
	Provider  string     `json:"provider"`
	Avatar    string     `json:"avatar"`
	SocialID  string     `json:"social_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
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
