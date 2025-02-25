package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func main() {
	a, b := 10, 0
	c, err := div(a, b)
	fmt.Println(c, err)

	f := test(100)
	// 函数作为返回值测试
	f()

	// 切片分组测试
	slice_test()

	map_test()
}

func test(x int) func() {
	return func() {
		println(x)
	}
}

func slice_test() {
	x := make([]int, 0, 10)
	for i := 0; i < 7; i++ {
		x = append(x, i)
	}

	fmt.Println(x)
}

func map_test() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	x, ok := m["b"]

	fmt.Println(x, ok)
	delete(m, "a")
	fmt.Println(m)

}
