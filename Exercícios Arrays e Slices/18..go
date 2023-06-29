package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	primes := make([]int, 0)
	count := 0
	num := 2

	for count < n {
		if isPrime(num) {
			primes = append(primes, num)
			count++
		}
		num++
	}

	fmt.Println("Array de números primos:", primes)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
