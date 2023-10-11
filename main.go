package main

import (
	"Go-Queue/queue"
	"fmt"
)

func main() {
	q := queue.New(4)
	q.Push("流沙", false)
	q.Push("爱很简单", false)
	q.Push("月亮代表谁的心", false)
	q.Push("普通朋友", false)
	for !q.Empty() {
		fmt.Println(q.Pop())
	}
	q.Push(1234, false)
	q.Push("思念，常常思念不常见面", false)
	q.Push("他怀疑Sam是虚拟的脸", false)
	q.Push(4.52, true)
	fmt.Println("------------------------------")
	for !q.Empty() {
		fmt.Println(q.Pop())
	}
}
