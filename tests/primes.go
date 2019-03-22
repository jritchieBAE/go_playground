package main

import "fmt"

func main() {
	gen := PrimeGenerator()
	for i := 0; i < 10000; i++ {
		gen()
	}
	fmt.Println(gen())
}

func PrimeGeneratorChannel() func() int {
	value := 2

	ch := make(chan int, 1)

	isPrime := func(num int) bool {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	next := func() int {
		for {
			if isPrime(value) {
				ch <- value
			}
			value = value + 1
		}
	}

	go next()

	return func() int {
		return <-ch
	}
}
