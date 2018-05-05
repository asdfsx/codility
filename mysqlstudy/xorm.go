package mysqlstudy

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"fmt"
)

type Xorm struct{
	Id int `xorm:"int(10) pk not null autoincr 'id'"`
	Name string `xorm:"varchar(25) not null unique 'name'"`
}

func CreateEngine() (*xorm.Engine, error) {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", MYSQLUSER, MYSQLPASSWORD, MYSQLADDR, MYSQLPORT, DATABASENAME)
	engine, err := xorm.NewEngine("mysql", connstr)
	engine.SetTableMapper(core.NewPrefixMapper(core.SameMapper{}, "prefix"))
	engine.SetColumnMapper(core.NewSuffixMapper(core.SameMapper{}, "suffix"))
	if err != nil{
		return nil, err
	}
	return engine, nil
}

func ORMCreateDatabase(engine *xorm.Engine) error {
	result, err := engine.Exec("create database if not exists xorm")
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func ORMDropDatabase(engine *xorm.Engine) error {
	result, err := engine.Exec("drop database if exists xorm")
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func ORMCreateTable(engine *xorm.Engine) error {
	err := engine.CreateTables(&Xorm{})
	if err != nil{
		return err
	}
	return nil
}

func ORMDropTable(engine *xorm.Engine) error {
	err := engine.DropTables(&Xorm{})
	if err != nil{
		return err
	}
	return nil
}

func ORMInsert(engine *xorm.Engine) error {
	for i:=0; i < 10 ; i++ {
		result, err := engine.Insert(&Xorm{Name:fmt.Sprintf("a_%02d", i)})
		if err != nil{
			return err
		}
		fmt.Println(result)
	}
	return nil
}

func ORMQuery(engine *xorm.Engine) error {
	cache := make([]Xorm, 0)
	err := engine.Find(&cache)
	if err != nil{
		return nil
	}

	for i:=0; i<len(cache); i++{
		fmt.Println(cache[i].Name)
	}
	return nil
}

func ORMUpdate(engine *xorm.Engine) error{
	cache := make([]Xorm, 0)
	err := engine.Find(&cache)
	if err != nil{
		return nil
	}

	for i:=0; i<len(cache); i++{
		cache[i].Name += "_update";
		result, err := engine.ID(cache[i].Id).Update(&cache[i])
		if err != nil{
			return nil
		}
		fmt.Println(result)
	}
	return nil
}