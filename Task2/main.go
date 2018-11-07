// Task2 project main.go
package main

import "fmt"

func main() {
	var a [4]string
	a[0] = "Artem"
	a[1] = "Belarus"
	a[2] = "246015"
	a[3] = "15"

	for i, v := range a {
		fmt.Printf("%d | %v\n", i, v)
	}
}
