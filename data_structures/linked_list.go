package main

import (
	"container/list"
	"fmt"
)

func pracLinkedList() {
	l := list.New()

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(5)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func main() {
	fmt.Println("linked list")
	pracLinkedList()
}
