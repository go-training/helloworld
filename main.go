package main

import "fmt"

func helloworld() string {
	return "Hello World!!"
}

func nihao() string {
	return "你好"
}

func main() {
	fmt.Println(helloworld())
	fmt.Println(nihao())
}
func add(a int, b int) int {
	return a + b
}
