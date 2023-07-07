package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	birthplace
}

type birthplace struct {
	city    string
	country string
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) changeCity(cityName string) {
	(*p).birthplace.city = cityName
}

func main() {
	player := person{
		firstName: "Declan",
		lastName:  "Rice",
		age:       24,
		birthplace: birthplace{
			city:    "Liverpool",
			country: "England",
		},
	}

	playerPointer := &player
	fmt.Println(&playerPointer)

	player.print()
	player.changeCity("London")
	fmt.Println("  ")
	player.print()
}
