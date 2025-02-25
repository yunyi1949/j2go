package main

// 多线程
import (
	"fmt"
	"strconv"
	"time"
)

func task(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d:%d\n", id, i)
		time.Sleep(time.Second)
	}
}

const (
	x = iota
	y
	z
)

func main() {
	go task(1)
	go task(2)
	time.Sleep(time.Second * 6)

	x, _ := strconv.Atoi("12")
	println(x)
}
