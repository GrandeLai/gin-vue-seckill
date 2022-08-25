package dao

import (
	"gin-vue-seckill/jdaw_product_srv/models"
)

func GetSeckillList(currentPage int, pageSize int) (list []*models.Seckills, err error, count int) {
	err = DB.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&list).Error
	seckill := []models.Seckills{}
	DB.Find(&seckill).Count(&count)
	return list, err, count
}

func GetProductList2() (list []*models.Products, err error) {
	err = DB.Find(&list).Error
	return list, err
}

func DeleteSeckill(id int) error {
	seckill := models.Seckills{Id: id}
	err := DB.Delete(&seckill).Error
	return err
}

func CreateSeckill(seckill *models.Seckills) error {
	err := DB.Create(seckill).Error
	return err
}

func GetSeckill(id int) (seckill models.Seckills, err error) {
	err = DB.Where("id = ?", id).First(&seckill).Error
	return seckill, err
}

func UpdateSeckill(seckill *models.Seckills) error {
	err := DB.Model(&seckill).Where("id=?", seckill.Id).Updates(&seckill).Error
	return err
}
