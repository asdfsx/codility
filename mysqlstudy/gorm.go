package mysqlstudy

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Gorm struct{
	Id int `gorm:"type:int 10;AUTO_INCREMENT;Column:prefix_id`
	Name string `gorm:"type:varchar(25);NOT NULL;UNIQUE;Column:name_sufix"`
}

func CreateDB()(*gorm.DB, error){
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MYSQLUSER, MYSQLPASSWORD, MYSQLADDR, MYSQLPORT, DATABASENAME)
	db, err := gorm.Open("mysql", connstr)
	if err != nil{
		return nil, err
	}
	return db, err
}

func GORMDropDatabase(db *gorm.DB) error {
	db = db.Exec("drop database if exists gorm")
	if db.Error != nil{
		return db.Error
	}
	return nil
}

func GORMDropTable(db *gorm.DB) error {
	db = db.DropTable(&Gorm{})
	if db.Error != nil{
		return db.Error
	}
	return nil
}

func GORMCreateDatabase(db *gorm.DB) error {
	db = db.Exec("create database if not exists gorm")
	if db.Error != nil{
		return db.Error
	}
	return nil
}

func GORMCreateTable(db *gorm.DB) error {
	db = db.AutoMigrate(&Gorm{})
	if db.Error != nil{
		return db.Error

	}
	return nil
}


func GORMInsert(db *gorm.DB) error {
	for i:=0; i < 10 ; i++ {
		db := db.Create(&Gorm{Name:fmt.Sprintf("a_%02d", i)})
		if db.Error != nil{
			return db.Error
		}
	}
	return nil
}

func GORMQuery(db *gorm.DB) error {
	cache := make([]Gorm, 0)
	db = db.Find(&cache)
	if db.Error != nil{
		return db.Error
	}

	for i:=0; i<len(cache); i++{
		fmt.Println(cache[i].Name)
	}
	return nil
}

func GORMUpdate(db *gorm.DB) error{
	cache := make([]Gorm, 0)
	db = db.Find(&cache)
	if db.Error != nil{
		return db.Error
	}

	for i:=0; i<len(cache); i++{
		cache[i].Name += "_update";
		db := db.Save(&cache[i])
		if db.Error != nil{
			return nil
		}
	}
	return nil
}