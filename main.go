package main

import "fmt"

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

	player.birthplace.city = "London"

	fmt.Printf("%+v", player)
}
