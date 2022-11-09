package mapx

import (
	"fmt"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestStruct2Map(t *testing.T) {
	convey.Convey("测试所有类型转map", t, func() {
		type TestStruct struct {
			S1     string    `json:"S1"`
			Int1   int       `json:"Int1"`
			Float1 float64   `json:"Float1"`
			Bool1  bool      `json:"Bool1"`
			T1     time.Time `json:"T1"`

			PtrS1     *string    `json:"PtrS1"`
			PtrInt1   *int       `json:"PtrInt1"`
			PtrFloat1 *float64   `json:"PtrFloat1"`
			PtrBool1  *bool      `json:"PtrBool1"`
			PtrTime   *time.Time `json:"PtrTime"`
		}

		ptrS1 := "s1"
		ptrInt1 := 1
		ptrFloat1 := 1.1
		ptrBool1 := true
		ptrTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")

		struc := TestStruct{
			S1:     "s1",
			Int1:   1,
			Float1: 1.1,
			Bool1:  true,
			T1:     ptrTime,

			PtrS1:     &ptrS1,
			PtrInt1:   &ptrInt1,
			PtrFloat1: &ptrFloat1,
			PtrBool1:  &ptrBool1,
			PtrTime:   &ptrTime,
		}

		m := Struct2Map(struc)
		fmt.Println(m)
		convey.So(m["S1"], convey.ShouldEqual, "s1")
		convey.So(m["Int1"], convey.ShouldEqual, 1)
		convey.So(m["Float1"], convey.ShouldEqual, 1.1)
		convey.So(m["Bool1"], convey.ShouldEqual, true)
		convey.So(m["T1"], convey.ShouldEqual, "2019-01-01 00:00:00")

		convey.So(m["PtrS1"], convey.ShouldEqual, "s1")
		convey.So(m["PtrInt1"], convey.ShouldEqual, 1)
		convey.So(m["PtrFloat1"], convey.ShouldEqual, 1.1)
		convey.So(m["PtrBool1"], convey.ShouldEqual, true)
		convey.So(m["PtrTime"], convey.ShouldEqual, "2019-01-01 00:00:00")
	})
}
