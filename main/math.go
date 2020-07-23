package main

import "fmt"

//一头母牛一年生一头小母牛，小母牛第二年生小母牛，12年后有多少母牛

func main() {
	a := 1

	for years := 1; years <= 12; years++ {
		if years == 1 {
			a = a + 1
		} else {
			a *= 2
		}
	}
	fmt.Printf("12年后有%d头母牛", a)
}
