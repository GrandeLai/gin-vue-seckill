package models

import "time"

type FrontUser struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password"`
	Desc       string    `gorm:"column:desc"` //用户描述
	Status     int       `gorm:"column:status"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (FrontUser) TableName() string {
	return "front_user"
}
