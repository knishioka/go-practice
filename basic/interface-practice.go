package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) speak() {
	fmt.Println(p.FirstName)
}

type human interface {
	speak()
}

func saySomething (h human) {
	h.speak()
}

func main() {
	p := Person{FirstName:"Test", LastName:"Test"}
	p.speak()
	saySomething(&p)
}
