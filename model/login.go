package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	gorm.Model
	Member_id   		int64  	`gorm:"type:int(10);unique_index;PRIMARY_KEY;AUTO_INCREMENT"`
	MemberNumber 		*string `gorm:"type:varchar(50)"` 
	MemberPassword		string	`gorm:"type:varchar(255)"`
	MemberPhone			string	`gorm:"type:varchar(11);unique;not null"`	// 设置手机号唯一并且不为空
	Created_at			int		`gorm:"type:int(10)"`
	Updated_at			int		`gorm:"type:int(10)"`
}

