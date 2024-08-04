package service

import (
	"context"
	"gin-temp/conf"
	"gin-temp/internal/global/datastore"
	"gin-temp/internal/global/errs"
	"gin-temp/internal/global/utils"
	"gin-temp/internal/model"
	"sync"
	"time"
)

type AccountService struct{}

var (
	accountService     *AccountService
	onceAccountService sync.Once
)

func GetAccountService() *AccountService {
	onceAccountService.Do(func() {
		accountService = &AccountService{}
	})
	return accountService
}

func (*AccountService) Login(account *model.Account) (string, error) {
	accountDao := model.GetAccountDao()
	user, err := accountDao.Login(account)
	if err != nil {
		return "", errs.NewServerErr("账号或密码错误", err)
	}
	expire := time.Duration(conf.Cfg.App.AuthExpire) * time.Second
	token, err := utils.GenerateAccessToken(user.ID, expire)
	if err != nil {
		return "", errs.NewServerErr("登录失败", err)
	}
	key := datastore.GetAccessKey(user.ID)
	err = datastore.Cache.Set(context.Background(), key, token, expire).Err()
	if err != nil {
		return "", errs.NewServerErr("登录时发生错误", err)
	}
	return token, nil
}

func (*AccountService) Register(user *model.User) error {
	accountDao := model.GetAccountDao()
	return accountDao.Register(user)
}
