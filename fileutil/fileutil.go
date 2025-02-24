package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountLines(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close() // 确保文件关闭

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount // 返回潜在扫描错误
}

func main() {
	fmt.Print(CountLines("README.md"))
}
