package goBenchExp

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"reflect"
	"unsafe"
)

func inlineFunc(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func emptyFunc(i int) {
	if i == 555 {
		fmt.Println(i)
	}
}

func simpleDefer(i int) (j int) {
	defer func() {
		j = i + 1
	}()
	return
}

func typeAssert() {
	var i interface{}
	i = 5
	if _, ok := i.(float64); ok {
		fmt.Println("float")
	}
}

func simpleReflectTypeOf() {
	var i interface{}
	i = float64(5)
	t := reflect.TypeOf(i)
	if t.String() == "int" {
		fmt.Println("size == 10")
	}
}

func simpleReflectValueOf() {
	var i interface{}
	i = float64(5)
	v := reflect.ValueOf(i)
	if v.String() == "int" {
		fmt.Println("type == int")
	}
}

func strToByteSlice(s string) []byte {
	return []byte(s)
}

func byteSliceToStr(b []byte) string {
	return string(b)
}

func byteSliceToStrDirect(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func getFromCh(ch chan string) {
	<-ch
}

func jsonUnmarshalStd(s []byte, v interface{}) interface{} {
	err := json.Unmarshal(s, v)
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func jsonUnmarshal3rd(s []byte, v interface{}) interface{} {
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(s, v)
	if err != nil {
		fmt.Println(err)
	}
	return v
}
