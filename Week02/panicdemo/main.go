package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		fmt.Println("Hello world!")
		panic("我来啦！")
	}()

	//go func() {
	//	defer func() {
	//		if err := recover(); err != nil{
	//			fmt.Println(err)
	//		}
	//	}()
	//	fmt.Println("Hello world!")
	//	panic("我来啦！")
	//}()

	//Go(func(){
	//	fmt.Println("Hello world!")
	//	panic("我来啦！")
	//})
	time.Sleep(5 * time.Second)
}

func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		x()
	}()
}
