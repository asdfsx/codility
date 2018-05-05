package testframework

import (
	"testing"
	. "github.com/asdfsx/codility/lesson1"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/prashantv/gostub"
	. "github.com/golang/mock/gomock"
)

func TestSolution(t *testing.T){
	Convey("expect result 5 ", t, func(){
		So(Solution(1041), ShouldEqual, 5)
	})


	stubs := Stub(&Fortest, "4test")
	defer stubs.Reset()
	Convey("should be 4test", t, func(){
		So(Fortest, ShouldEqual, "4test")
	})

	var Solution2 = Solution
	stubs1 := Stub(&Solution2, func(N int) int {
		return 0
	})
	defer stubs1.Reset()
	Convey("should return 0", t, func(){
		So(Solution2(111), ShouldEqual, 0)
	})

	var Solution3 = Solution
	stubs2 := StubFunc(&Solution3, 0)
	defer stubs2.Reset()
	Convey("should return 0", t, func(){
		So(Solution2(111), ShouldEqual, 0)
	})

	var newRedisRepo2 = newRedisRepo
	Convey("test obj demo", t, func() {
		Convey("create obj", func() {
			ctrl := NewController(t)
			defer ctrl.Finish()
			mockRepo := NewMockRepository(ctrl)
			mockRepo.EXPECT().Retrieve(Any()).Return(nil, nil)
			mockRepo.EXPECT().Create(Any(), Any()).Return(nil)
			mockRepo.EXPECT().Retrieve(Any()).Return(nil, nil)
			stubs := StubFunc(&newRedisRepo2, mockRepo)
			defer stubs.Reset()
			So(repo.Create("", nil), ShouldBeNil)
		})
	})
}