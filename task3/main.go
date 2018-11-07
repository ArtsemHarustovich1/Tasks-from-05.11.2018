// task3 project main.go
package main

import "fmt"

type Person struct {
	Name    string
	Country string
	Zip     string
	Id      int
}

func (p Person) String() string {
	return fmt.Sprintf("%v%v%v%v", p.Name, p.Country, p.Zip, p.Id)
}

func main() {
	a := Person{"Artem\n", "Belarus\n", "246015\n", 15}
	fmt.Println(a)
}
