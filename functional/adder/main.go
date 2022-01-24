package main

import "fmt"

func adder() func(i int) int {
	var sum = 0
	return func(v int) int {
		sum = sum + v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	fmt.Println(base)
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	// var add = adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("0 + 1 + ... %d = %d \n", i, add(i))
	// }
	var a = adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... %d = %d \n", i, s)
	}
}
