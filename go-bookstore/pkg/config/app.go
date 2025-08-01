package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "Username:Password/simplerest?charset=utf8&parseTimeTrue&loc=Local")
	if err != nil {
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}