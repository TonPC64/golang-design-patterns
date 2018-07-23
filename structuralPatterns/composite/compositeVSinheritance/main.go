package main

import "fmt"

type Parent struct {
	SomeField int
}

type Son struct {
	Parent
}

func GetParentField(p Parent) {
	fmt.Println(p.SomeField)
}

func main() {
	son := Son{}
	GetParentField(son.Parent)
}
