package gorm_related

import "github.com/jinzhu/gorm"

type User struct {
	// 里面含有一些默认的字段
	gorm.Model
	Age  int    `gorm:"type:int;default:18"`
	Name string `gorm:"type:varchar(20);not null"`
	Sex  string `gorm:"type:varchar(6)"`
}
