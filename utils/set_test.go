package utils

import (
	"flag"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var stringSet *StringSet

func TestMain(m *testing.M) {

	stringSet = NewStringSet()
	flag.Parse()
	exitCode := m.Run()

	// 退出
	os.Exit(exitCode)
}

func TestStringSet_Add(t *testing.T) {
	Convey("check", t, func() {
		stringSet.Add("12241")
		So(stringSet, ShouldNotBeNil)
	})
}
func TestStringSet_Has(t *testing.T) {
	Convey("has", t, func() {
		So(stringSet.Has("12241"), ShouldBeTrue)
	})
}
func TestStringSet_Remove(t *testing.T) {
	Convey("Remove", t, func() {
		stringSet.Add("ssml")
		So(stringSet.Has("ssml"), ShouldBeTrue)
		stringSet.Remove("ssml")
		So(stringSet.Has("ssml"), ShouldBeFalse)
	})
}
func TestStringSet_Len(t *testing.T) {
	Convey("Len", t, func() {
		So(stringSet.Len(), ShouldEqual, 1)
		So(stringSet.IsEmpty(), ShouldBeFalse)
	})
}
func TestStringSet_Clear(t *testing.T) {
	Convey("clear", t, func() {
		stringSet.Clear()
		So(stringSet.Len(), ShouldEqual, 0)
	})
}

func TestStringSet_IsEmpty(t *testing.T) {
	Convey("siempty", t, func() {
		stringSet.Clear()
		So(stringSet.IsEmpty(), ShouldBeTrue)
	})
}

func TestIntSet(t *testing.T) {
	intSet := NewIntSet()
	Convey("add", t, func() {
		intSet.Add(1)
		So(intSet.Len(), ShouldEqual, 1)
		intSet.Add(1)
		So(intSet.Len(), ShouldEqual, 1)
		intSet.Add(3)
		So(intSet.Len(), ShouldEqual, 2)
	})

	Convey("remove", t, func() {
		intSet.Remove(1)
		So(intSet.Len(), ShouldEqual, 1)
		So(intSet.Has(1), ShouldBeFalse)
		So(intSet.IsEmpty(), ShouldBeFalse)
	})

	Convey("clear", t, func() {
		intSet.Clear()
		So(intSet.Len(), ShouldEqual, 0)
		So(intSet.Has(1), ShouldBeFalse)
		So(intSet.IsEmpty(), ShouldBeTrue)
	})
}
