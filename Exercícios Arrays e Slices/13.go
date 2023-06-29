package main

import "fmt"

func main() {
	array := [7]int{2, 4, 6, 8, 10, 12, 14}

	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	array[0] += num
	array[len(array)-1] += num

	fmt.Println("Array resultante:", array)
}
