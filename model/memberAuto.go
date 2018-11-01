package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

//隧道权限管理
type MemberAuto struct {
	gorm.Model
	Member_id   	int64  	`gorm:"type:int(10);unique_index;PRIMARY_KEY;"`
	Token		 	*string `gorm:"unique;not null"`
	Permission		string	`gorm:"type:varchar(1024)"`
	Updated_at     	int  	`gorm:"type:int(10)"`
}
  