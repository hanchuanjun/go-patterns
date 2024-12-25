package main

import "fmt"

type Language interface {
	say()
}

type person struct {
	name string
	country string
	language Language
}


func (p *person) say() {
	p.language.say()
}

type chinese struct {	
	name string
}

func (c *chinese) say() {
	fmt.Println("chinese say")
}

type english struct {
}

func (e *english) say() {
	fmt.Println("english say")
}



func main() {
	p := &person{
		name:     "John",
		country:  "UK",
		language: &english{},
	}
	p.say()
}

