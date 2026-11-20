package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	left := num / 1000
	right := num % 1000

	sumLeft := (left / 100) + ((left / 10) % 10) + (left % 10)
	sumRight := (right / 100) + ((right / 10) % 10) + (right % 10)

	if num < 99999 {
		fmt.Println("Число меньше шестизначного")
		return
	}

	if sumLeft == sumRight {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
