package main

import "fmt"

func main() {
	array := [5]int{7, 3, 12, 9, 5}

	slice := make([]int, 0)

	for _, num := range array {
		if num%3 == 0 {
			slice = append(slice, num)
		}
	}

	fmt.Println("Novo Slice:", slice)
}
