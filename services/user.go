package services

import (
	"errors"
	"fmt"
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
func DelUser(id uint64) error {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	return model.DB.Delete(&user, id).Error
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

func UpdateUser(id uint64, nickname string, signature string) error {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	if nickname != "" {
		user.Nickname = nickname
	}
	if signature != "" {
		user.Signature = signature
	}
	return model.DB.Save(user).Error
}

func UpdatePassword(id uint64, password string) error {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	return model.DB.Save(user).Error
}

func GetAllUsers(pageSize int, page int) ([]model.User, int64, error) {
	//处理错误的页码
	if page <= 0 {
		page = 1
	}

	var users []model.User
	var count int64
	model.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)
	model.DB.Model(&model.User{}).Count(&count)
	return users, count, nil
}

func SearchUsers(name string, pageSize int, page int) ([]model.User, int64, error) {
	var users []model.User
	name = fmt.Sprintf("%%%s%%", name)
	var count int64
	db := model.DB.Model(&model.User{}).Where("username Like ?", name).Or("nickname Like ?", name).Count(&count)
	if page < 1 {
		page = 1
	}
	if count > (int64)((page-1)*pageSize) {
		db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)
	} else {
		db.Find(&users)
	}
	return users, count, nil
}

func SetAdminUser(userid uint64) error {
	user := model.User{ID: userid}
	return model.DB.First(&user).Update("isAdmin", 1).Error
}

func RemoveAdmin(userid uint64) error {
	user := model.User{ID: userid}
	return model.DB.First(&user).Update("isAdmin", 0).Error
}
