package etcd

import (
	"testing"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/coreos/etcd/client"
	"time"
	"fmt"
)

func TestEtcd(t *testing.T){
	stubs := New()
	stubs.Stub(&Endpoints, []string{"http://127.0.0.1:2379"})
	defer stubs.Reset()

	var etcdCli client.Client
	var err error

	Convey("Connect", t, func(){
		etcdCli, err = Connect()
		So(err, ShouldBeNil)
	})

	Convey("Clean", t, func(){
		isExist, err := Exist(etcdCli, "/test")
		fmt.Println("----", isExist, err)
		So(err, ShouldBeNil)
		if isExist {
		    err := Delete(etcdCli, "/test", &client.DeleteOptions{Dir:true, Recursive:true})
		    So(err, ShouldBeNil)
		}
	})

	Convey("Create", t, func(){
		err := Create(etcdCli, "/test", "ttt")
		So(err, ShouldBeNil)
	})

	Convey("Get", t, func(){
		result, err := Get(etcdCli, "/test", &client.GetOptions{})
		t.Log(result)
		So(err, ShouldBeNil)
	})

	Convey("Set", t, func(){
		result, err := Set(etcdCli, "/test/test1", "t1", &client.SetOptions{})
		t.Log(result)
		So(err, ShouldNotBeNil)
	})

	Convey("Delete", t, func(){
		err := Delete(etcdCli, "/test", &client.DeleteOptions{})
		t.Log(err)
		So(err, ShouldBeNil)
	})

	Convey("Set Dir", t, func(){
		result, err := Set(etcdCli, "/test", "dir",  &client.SetOptions{PrevExist: client.PrevIgnore, Dir:true})
		t.Log(result)
		So(err, ShouldBeNil)
	})

	Convey("Set /test/test1", t, func(){
		result, err := Set(etcdCli, "/test/test1", "test1", &client.SetOptions{})
		t.Log(result)
		So(err, ShouldBeNil)
	})

	Convey("Get /test/test1", t, func(){
		result, err := Get(etcdCli, "/test/test1", &client.GetOptions{})
		t.Log(result)
		So(err, ShouldBeNil)
	})

	Convey("Set TTL /test/test2", t, func(){
		result, err := Set(etcdCli, "/test/test2", "test2", &client.SetOptions{TTL:2 * time.Second})
		t.Log(result)
		So(err, ShouldBeNil)
	})

	Convey("Get TTL /test/test2", t, func(){
		isExist, err := Exist(etcdCli, "/test/test2")
		t.Log("before", isExist)
		So(isExist, ShouldBeTrue)
		So(err, ShouldBeNil)

		time.Sleep(6 * time.Second)

		isExist, err = Exist(etcdCli, "/test/test2")
		t.Log("after", isExist)
		So(isExist, ShouldBeFalse)
		So(err, ShouldNotBeNil)
	})

	Convey("Set TTL /test/test3", t, func(){
		result, err := Set(etcdCli, "/test/test3", "test3", &client.SetOptions{TTL:2 * time.Second})
		t.Log(result)
		So(err, ShouldBeNil)
	})

	Convey("WATCH TTL /test/test3", t, func(){
		isExist, err := Exist(etcdCli, "/test/test3")
		t.Log("before", isExist)
		So(isExist, ShouldBeTrue)
		So(err, ShouldBeNil)

		resultChan1 := Watch(etcdCli, "/test", &client.WatcherOptions{Recursive:true})
		resultChan2 := Watch(etcdCli, "/test/test3", &client.WatcherOptions{})

		time.Sleep(6 * time.Second)


		result, _ := <- resultChan1
		So(result, ShouldEqual, "expire")

		result2, _ := <- resultChan2
		So(result2, ShouldEqual, "expire")

		isExist, err = Exist(etcdCli, "/test/test2")
		t.Log("after", isExist)
		So(isExist, ShouldBeFalse)
		So(err, ShouldNotBeNil)
	})
}