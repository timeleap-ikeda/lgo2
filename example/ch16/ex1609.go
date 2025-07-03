package main

import (
	"fmt"
	"reflect"
	"time"
)


func MakeTimedFunction(f any) any { //liststart1
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fv.Call(in)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}  //listend1

func timeMe(a int) int {  //liststart2
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}

func main() {
	timed := MakeTimedFunction(timeMe).(func(int) int)
	fmt.Println(timed(2))
}  //listend2
