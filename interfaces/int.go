package main

import (
	"bufio"
	"fmt"
	"os"
)

type StringProvider interface {
	GiveMeAString() string
}

type DefinedString struct {
	value string
}

func (d DefinedString) GiveMeAString() string {
	return d.value
}

type StringRequester struct {
}

func (s StringRequester) GiveMeAString() string {
	fmt.Printf("Please enter a value: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	stringer1 := DefinedString{"hello"}
	stringer2 := StringRequester{}
	Print(stringer1)
	Print(stringer2)
}

func Print(s StringProvider) {
	fmt.Println(s.GiveMeAString())
}
