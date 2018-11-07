// main project main.go
package main

import (
	"fmt"
)

type Human struct {
	Name string
	Info struct {
		Country string
		Zip     string
		Id      int
	}
}

func id(x, y int) int {
	return x * y
}

func main() {
	c := Human{
		Name: "Artem",
	}
	c.Info.Country = `Belarus`
	c.Info.Zip = `246015`
	c.Info.Id = id(3, 5)
	fmt.Println(c)
}
