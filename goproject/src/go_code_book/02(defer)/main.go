package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0

	defer c()

	return

}

func c() (i int) {
	fmt.Println("-------")
	defer func() { i++ }()
	return 1
}
