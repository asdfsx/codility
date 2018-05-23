package mysqlstudy

import (
	"testing"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
)

func TestGorm(t *testing.T) {
	stubs := New()
	stubs.Stub(&MYSQLUSER, "root")
	stubs.Stub(&MYSQLPASSWORD, "root")
	stubs.Stub(&MYSQLADDR, "192.168.99.100")
	stubs.Stub(&MYSQLPORT, 3306)
	stubs.Stub(&DATABASENAME, "mysql")
	defer stubs.Reset()

	var db *gorm.DB
	var err error

	Convey("create engine", t, func(){
		db, err = CreateDB()
		So(err, ShouldEqual, nil)
	})
	Convey("create database", t, func(){
		So(GORMCreateDatabase(db), ShouldEqual, nil)
	})
	stubs.Stub(&DATABASENAME, "gorm")
	Convey("create new db", t, func(){
		db, err = CreateDB()
		So(err, ShouldEqual, nil)
	})
	Convey("create table", t, func(){
		So(GORMCreateTable(db), ShouldEqual, nil)
	})
	Convey("insert", t, func(){
		So(GORMInsert(db), ShouldEqual, nil)
	})
	Convey("query", t, func(){
		So(GORMQuery(db), ShouldEqual, nil)
	})
	Convey("update", t, func(){
		So(GORMUpdate(db), ShouldEqual, nil)
	})

	Convey("query", t, func(){
		So(GORMQuery(db), ShouldEqual, nil)
	})

	Convey("drop table", t, func(){
		So(GORMDropTable(db), ShouldEqual, nil)
	})

	Convey("drop database", t, func(){
		So(GORMDropDatabase(db), ShouldEqual, nil)
	})
}