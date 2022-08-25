package models

import "time"

type AdminUser struct {
	Id         int       `gorm:"column:id"`
	UserId     string    `gorm:"column:user_id"`
	UserName   string    `gorm:"column:user_name"`
	Password   string    `gorm:"column:password"`
	Desc       string    `gorm:"column:desc"`
	Status     int       `gorm:"column:status"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (AdminUser) TableName() string {
	return "admin_user"

}
