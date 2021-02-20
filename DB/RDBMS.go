package DB

import "github.com/jinzhu/gorm"

type RDBMS struct {
	db *gorm.DB
}

func (r *RDBMS) Control(do func(db *gorm.DB) *gorm.DB,result []interface{}) error {
	return do(r.db).Error
}


func (r *RDBMS) CloseDB() error {
	return r.db.Close()
}

func ConnectDB() (*gorm.DB,error) {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "password"
	PROTOCOL := "tcp(zabuton_db:3306)"
	DBNAME   := "zabuton"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	return gorm.Open(DBMS, CONNECT)
}
