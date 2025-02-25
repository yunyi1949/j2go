package main

import "fmt"

func main() {
	var p *int

	i := 10

	p = &i
	fmt.Println(p)
	fmt.Println(*p)

	*p = 100 // 通过指针修改对应地址的值
	fmt.Println(*p)

}
