package conn
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
var MYSqlConns *gorm.DB

func MYSqlConn() (*gorm.DB) {
	MYSqlConns, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer MYSqlConns.Close()
	if err != nil {
		panic("failed to connect database")
	}
	return MYSqlConns
}
