package services

import (
	"errors"
	"github.com/96368a/LuoYiMusic-server-api/model"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserName(username string) bool {
	var user model.User
	err := model.DB.First(&user, "username = ?", username).Error

	return err != nil
}

func AddUser(nickname string, username string, password string) (*model.User, error) {
	if ok := CheckUserName(username); ok != true {
		return nil, errors.New("用户已存在")
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := model.User{
		Nickname:  nickname,
		Username:  username,
		Password:  string(hashPassword),
		Signature: "还没有签名呢...",
		IsAdmin:   0,
	}
	err := model.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Login(username string, password string) (*model.User, error) {
	var user model.User
	err := model.DB.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

func UpdateUser(user *model.User) error {
	return model.DB.Save(user).Error
}
