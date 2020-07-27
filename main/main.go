package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type Struct struct{}

func (s *Struct) Call(p interface{}) {
	fmt.Println("From struct info", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}

func main() {
	// 声明接口变量
	var invoker Invoker
	// 实例化结构体
	s := new(Struct)
	// 将实例化的结构体赋值到接口
	invoker = s
	// 使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("hello")
	// 将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	invoker.Call("hello")
}
