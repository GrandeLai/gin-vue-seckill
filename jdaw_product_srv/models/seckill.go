package models

import "time"

type Seckills struct {
	Id         int       `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Price      float32   `gorm:"column:price"`
	Num        int       `gorm:"column:num"`
	PId        string    `gorm:"column:pid"`        //商品外键
	StartTime  time.Time `gorm:"column:start_time"` //活动开始时间
	EndTime    time.Time `gorm:"column:end_time"`   //活动结束时间
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Seckills) TableName() string {
	return "product_seckill"
}
