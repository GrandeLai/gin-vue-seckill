package models

import "time"

type Products struct {
	Id         int       `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Price      float32   `gorm:"column:price"`
	Num        int       `gorm:"column:num"`
	Unit       string    `gorm:"column:unit"` //商品单位
	Pic        string    `gorm:"column:pic"`  //图片
	Desc       string    `gorm:"column:desc"`
	CreateTime time.Time `gorm:"column:create_time"`

	Seckills []Seckills `gorm:"ForeignKey:PId;AssociationForeignKey:Id"`
}

func (Products) TableName() string {
	return "product"
}
