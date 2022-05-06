package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 发送 sum 到通道 c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	// 并发
	// go say("world")
	// say("hello")

	// 通道
	// s := []int{1, 2, 3, 4, 5, 6}
	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c // 从通道 c 中读取
	// fmt.Println(x, y, x+y)

	// 通道缓冲区
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// 遍历和关闭通道
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// c 发送完 10 个数据之后就关闭通道
	for i := range c {
		fmt.Println(i)
	}
}
