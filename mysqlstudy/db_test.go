package mysqlstudy

import (
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/prashantv/gostub"
	"database/sql"
	"testing"
	"fmt"
)



func TestDB1(t *testing.T){
	// create connection
	connstr := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "root", "192.168.99.100:3306", "mysql")
	dbConn, err := sql.Open("mysql", connstr)
	if err != nil{
		t.Error(err)
	}

	// create database
	err = CreateDatabase(dbConn)
	if err != nil {
		t.Error(err)
	}
	// create table
	err = CreateTable(dbConn)
	if err != nil {
		t.Error(err)
	}
	// insert
	err = Insert(dbConn)
	if err != nil {
		t.Error(err)
	}
	// select
	err = Query(dbConn)
	if err != nil {
		t.Error(err)
	}
	// update
	err = Update(dbConn)
	if err != nil {
		t.Error(err)
	}
	// drop table
	err = DropTable(dbConn)
	if err != nil {
		t.Error(err)
	}
	// drop database
	err = DropDatabase(dbConn)
	if err != nil {
		t.Error(err)
	}
}

func TestDB2(t *testing.T){
	userstub := Stub(&MYSQLUSER, "root")
	defer userstub.Reset()
	pwdstub := Stub(&MYSQLPASSWORD, "root")
	defer pwdstub.Reset()
	addrstub := Stub(&MYSQLADDR, "192.168.99.100")
	defer addrstub.Reset()
	portstub := Stub(&MYSQLPORT, 3306)
	defer portstub.Reset()
	dbnamestub := Stub(&DATABASENAME, "mysql")
	defer dbnamestub.Reset()

	var dbConn *sql.DB
	var err error

	Convey("CreateConnection", t, func(){
		dbConn, err = CreateConnection()
		So(err, ShouldEqual, nil)
	})

	Convey("CreateDatabase", t, func(){
		err = CreateDatabase(dbConn)
		So(err, ShouldEqual, nil)
	})

	Convey("CreateTable", t, func(){
		err = CreateTable(dbConn)
		So(err, ShouldEqual, nil)
	})

	Convey("Insert", t, func(){
		err = Insert(dbConn)
		So(err, ShouldEqual, nil)
	})

	Convey("Update", t, func(){
		err = Update(dbConn)
		So(err, ShouldEqual, nil)
	})

	Convey("DropTable", t, func(){
		err = DropTable(dbConn)
		So(err, ShouldEqual, nil)
	})

	Convey("DropDatabase", t, func(){
		err = DropDatabase(dbConn)
		So(err, ShouldEqual, nil)
	})

}
