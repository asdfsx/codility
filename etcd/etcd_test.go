package etcd

import (
	"testing"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/coreos/etcd/client"
)

func TestEtcd(t *testing.T){
	stubs := New()
	stubs.Stub(&Endpoints, []string{"http://192.168.99.100:2379"})
	defer stubs.Reset()

	var etcdCli client.Client
	var err error

	Convey("Connect", t, func(){
		etcdCli, err = Connect()
		So(err, ShouldBeNil)
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
}