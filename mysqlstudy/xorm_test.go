package mysqlstudy

import (
	"testing"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/go-xorm/xorm"
)

func TestXorm(t *testing.T){
	stubs := New()
	stubs.Stub(&MYSQLUSER, "root")
	stubs.Stub(&MYSQLPASSWORD, "root")
	stubs.Stub(&MYSQLADDR, "192.168.99.100")
	stubs.Stub(&MYSQLPORT, 3306)
	stubs.Stub(&DATABASENAME, "mysql")
	defer stubs.Reset()

	var engine *xorm.Engine
	var err error
	Convey("create engine", t, func(){
		engine, err = CreateEngine()
		So(err, ShouldEqual, nil)
	})
	Convey("create database", t, func(){
		So(ORMCreateDatabase(engine), ShouldEqual, nil)
	})
	stubs.Stub(&DATABASENAME, "xorm")
	Convey("create new engine", t, func(){
		engine, err = CreateEngine()
		So(err, ShouldEqual, nil)
	})
	Convey("create table", t, func(){
		So(ORMCreateTable(engine), ShouldEqual, nil)
	})

	Convey("insert", t, func(){
		So(ORMInsert(engine), ShouldEqual, nil)
	})

	Convey("query", t, func(){
		So(ORMQuery(engine), ShouldEqual, nil)
	})

	Convey("update", t, func(){
		So(ORMUpdate(engine), ShouldEqual, nil)
	})

	Convey("query", t, func(){
		So(ORMQuery(engine), ShouldEqual, nil)
	})

	Convey("drop table", t, func(){
		So(ORMDropTable(engine), ShouldEqual, nil)
	})

	Convey("drop database", t, func(){
		So(ORMDropDatabase(engine), ShouldEqual, nil)
	})
}