package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	oops()
}

func oops() {
	panic("oh no!")
}
