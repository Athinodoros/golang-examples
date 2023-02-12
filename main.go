package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int16
	email string
}

func (p *Person) addAge() {
	p.age += 1
}

func main() {

	person1 := Person{name: "nos", age: 33, email: "athinodoros.sgouromallis@gmail.com"}
	person1.addAge()
	fmt.Println(person1)

}
