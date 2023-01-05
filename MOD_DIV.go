package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)

	fmt.Println(a, "div", b, "=", a/b)
	fmt.Println(a, "mod", b, "=", a%b)
}
