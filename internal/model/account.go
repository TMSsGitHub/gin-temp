package model

import (
	"gin-temp/internal/global/datastore"
	"sync"
)

type Account struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountDao struct{}

var (
	accountDao  *AccountDao
	onceAccount sync.Once
)

func GetAccountDao() *AccountDao {
	onceAccount.Do(func() {
		accountDao = &AccountDao{}
	})
	return accountDao
}

func (*AccountDao) Login(account *Account) (*User, error) {
	var user User
	res := datastore.DB.
		Where("phone = ?", account.Phone).
		Where("password = ?", account.Password).
		Where("deleted_at = 0").
		First(&user)
	return &user, res.Error
}

func (*AccountDao) Register(user *User) error {
	res := datastore.DB.Create(&user)
	return res.Error
}
