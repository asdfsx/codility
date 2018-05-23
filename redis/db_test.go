package redis1

import (
	"testing"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/gomodule/redigo/redis"
	"time"
)

func TestRedis(t *testing.T){
	stubs := New()
	stubs.Stub(&NETWORK, "tcp")
	stubs.Stub(&ADDR, "192.168.99.100")
	stubs.Stub(&PORT, 6379)
	defer stubs.Reset()

	var conn redis.Conn
	var err error

	Convey("Connect", t, func(){
		conn, err = Connect()
		So(err, ShouldBeNil)
	})

	Convey("Set", t, func(){
		So(Set(conn, "test", "test"), ShouldBeNil)
	})

	Convey("Get", t, func(){
		result, err := Get(conn, "test")
		So(result, ShouldEqual, "test")
		So(err, ShouldBeNil)
	})

	Convey("Del", t, func(){
		So(Del(conn, "test"), ShouldBeNil)
		result, err := Exists(conn, "test")
		So(err, ShouldBeNil)
		So(result, ShouldBeFalse)
	})

	Convey("TTL", t, func(){
		So(TTL(conn, "test", "test", 5), ShouldBeNil)
		isExists, err := Exists(conn, "test")
		So(err, ShouldBeNil)
		So(isExists, ShouldBeTrue)
		time.Sleep(time.Second * 5)
		isExists, err = Exists(conn, "test")
		So(err, ShouldBeNil)
		So(isExists, ShouldBeFalse)
	})

	Convey("LPUSH", t, func(){
		So(Lpush(conn, "test","1"), ShouldBeNil)
		So(Lpush(conn, "test","2"), ShouldBeNil)
		So(Lpush(conn, "test","3"), ShouldBeNil)
		length, err := Llen(conn, "test")
		So(err, ShouldBeNil)
		So(length, ShouldEqual,3)
	})

	Convey("LPOP", t, func(){
		result, err := Lpop(conn, "test")
		So(err, ShouldBeNil)
		So(result, ShouldEqual,"3")

		result, err = Lpop(conn, "test")
		So(err, ShouldBeNil)
		So(result, ShouldEqual,"2")

		result, err = Lpop(conn, "test")
		So(err, ShouldBeNil)
		So(result, ShouldEqual,"1")

		result, err = Lpop(conn, "test")
		t.Log("----", err)
		So(err, ShouldNotBeNil)
		So(result, ShouldEqual, "")
	})
}