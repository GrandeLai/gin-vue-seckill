package user

import (
	mysql "gin-vue-seckill/jdaw_user_srv/data_source"
	"gin-vue-seckill/jdaw_user_srv/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = mysql.Db

func CreateFrontUser(user *models.FrontUser) (err error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hasedPassword)
	err = DB.Create(user).Error
	return err
}

func GetFrontUser(mail string) (user *models.FrontUser, err error) {
	user = new(models.FrontUser)
	err = DB.Where("email=?", mail).First(user).Error
	return user, err
}
