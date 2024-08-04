package model

import (
	"gin-temp/internal/global/datastore"
	"gin-temp/internal/global/utils"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Sex       string `json:"sex,omitempty"`
	BirthAt   uint64 `json:"birth_at,omitempty"`
	BaseModel
}

func (user *User) BeforeCreate(_ *gorm.DB) error {
	user.ID = utils.GetSnowflakeID()
	return nil
}

type UserDao struct{}

var (
	userDao     *UserDao
	onceUserDao sync.Once
)

func GetUserDao() *UserDao {
	onceUserDao.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (userDao *UserDao) GetByPhone(phone string) (*User, error) {
	var user User
	res := datastore.DB.
		Where("phone = ?", phone).
		Where("deleted_at = 0").
		Find(&user)
	return &user, res.Error
}
