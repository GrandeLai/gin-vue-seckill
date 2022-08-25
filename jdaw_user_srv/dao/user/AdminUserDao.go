package user

import "gin-vue-seckill/jdaw_user_srv/models"

func GetAdminUser(username string) (user *models.AdminUser, err error) {
	user = new(models.AdminUser)
	err = DB.Where("user_name=?", username).First(user).Error
	return user, err
}

func GetFrontUserList(currentPage int, pageSize int) (list []*models.FrontUser, err error, count int) {
	err = DB.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&list).Error
	users_count := []models.FrontUser{}
	DB.Find(&users_count).Count(&count)
	return list, err, count
}
