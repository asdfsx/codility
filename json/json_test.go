package json

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func TestJson(t *testing.T){
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	var b []byte
	var err error

	Convey("marshal test", t, func(){
		b, err = Marshal(group)
		So(err, ShouldBeNil)
		t.Log(b)
	})

	Convey("unmarshal test", t, func(){
		var tmp ColorGroup
		err = Unmarshal(b, &tmp)
		So(err, ShouldBeNil)
		t.Log(tmp)
	})

}