package SqlliteDb

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db = new(gorm.DB)

var db_list = make(map[string]*gorm.DB)

// openDB 打开数据库
func OpenDB() *gorm.DB {
	//打开数据库，如果不存在，则创建
	db, err := gorm.Open(sqlite.Open("./price.db"), &gorm.Config{})
	fmt.Println(err)
	return db
}

// openDB 打开数据库
func OpenTableDB(tableName string) *gorm.DB {
	_, ok := db_list[tableName]
	if ok {
		return db_list[tableName]
	} else {
		db_list[tableName], _ = gorm.Open(sqlite.Open("./"+tableName+".db"), &gorm.Config{})
		return db_list[tableName]
	}
	//打开数据库，如果不存在，则创建
}
