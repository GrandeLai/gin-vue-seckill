package dao

import (
	mysql "gin-vue-seckill/jdaw_product_srv/data_source"
	"gin-vue-seckill/jdaw_product_srv/models"
)

var DB = mysql.Db

func GetProductList(currentPage int, pageSize int) (list []*models.Products, err error, count int) {
	err = DB.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&list).Error
	product := []models.Products{}
	DB.Find(&product).Count(&count)
	return list, err, count
}

func CreateProduct(product *models.Products) error {
	err := DB.Create(product).Error
	return err
}

func DeleteProduct(id int) error {
	product := models.Products{Id: id}
	err := DB.Delete(&product).Error
	return err
}

func GetProduct(id int) (models.Products, error) {
	product := models.Products{}
	err := DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func UpdateProduct(product *models.Products) error {
	err := DB.Model(&product).Where("id=?", product.Id).Updates(&product).Error
	return err
}
