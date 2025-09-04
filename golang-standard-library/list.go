package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("kubbil")
	data.PushBack("nur")
	data.PushBack("iman")
	data.PushBack("kubil")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // kubbil

	next := head.Next()
	fmt.Println(next.Value) // nur

	next = next.Next()
	fmt.Println(next.Value) // iman

	next = next.Next()
	fmt.Println(next.Value) // kubil

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
