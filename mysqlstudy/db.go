package mysqlstudy

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var (
	MYSQLUSER string
	MYSQLPASSWORD string
	MYSQLADDR string
	MYSQLPORT int
	DATABASENAME string
)

func CreateConnection() (*sql.DB, error) {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", MYSQLUSER, MYSQLPASSWORD, MYSQLADDR, MYSQLPORT, DATABASENAME)
	dbConn, err := sql.Open("mysql", connstr)
	if err != nil{
		return nil, err
	}
	return dbConn, nil
}

func CreateDatabase(db *sql.DB) error {
	sql := "create database sxtest"
	result, err := db.Exec(sql)
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func DropDatabase(db *sql.DB) error {
	sql := "drop database sxtest"
	result, err := db.Exec(sql)
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func CreateTable(db *sql.DB) error {
	sql := "create table sxtest.sxtest(id int)";
	result, err := db.Exec(sql)
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func DropTable(db *sql.DB) error {
	sql := "drop table sxtest.sxtest";
	result, err := db.Exec(sql)
	if err != nil{
		return err
	}
	fmt.Println(result)
	return nil
}

func Insert(db *sql.DB) error {
	sql := "insert into sxtest.sxtest (id) values (?)"
	stmt, err := db.Prepare(sql)
	if err != nil{
		return err
	}
	for i:=0; i<10;i++{
		result, err := stmt.Exec(i)
		if err != nil {
			return err
		}
		fmt.Println(result)
	}
	return nil
}

func Update(db *sql.DB) error {
	sql := "update sxtest.sxtest set id=id+10"
	result, err := db.Exec(sql)
	if err != nil{
		return nil
	}
	fmt.Println(result)
	return nil
}

func Query(db *sql.DB) error {
	sql := "select * from sxtest.sxtest"
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	for rows.Next() {
		fmt.Println(rows.Columns())
	}
	return nil
}