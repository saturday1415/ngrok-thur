package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	gorm.Model
	Member_id   	int64  `gorm:"type:int(10);unique_index;PRIMARY_KEY;AUTO_INCREMENT"`
	MemberNumber 	*string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Password		string	`gorm:"type:varchar(255);"`
	Email       	string  `gorm:"type:varchar(100);unique_index"`
}