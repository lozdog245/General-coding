package main

import (
	"fmt"
	"log"
)

type netflixshows struct {
	title string

	rating float32
}

type dcactionheros struct {
	name string

	age float32

	colour string

	address string
}

type boatrace struct {
	name string

	number int

	race string

	winner winner
}

type winner struct {
	place int

	time float64
}

func main() {

	h := boatrace{
		name:   "Fierce Sailing",
		number: 5,
		race:   "Rapid Race",
		winner: winner{
			place: 1,
			time:  0.987,
		},
	}
	fmt.Printf("%+v", h)

	var s dcactionheros

	s.name = "Superman"
	s.age = 34
	s.colour = "red and white"
	s.address = "4 Superman Avenue, Superville"
	fmt.Printf("%+v", s)

	var n netflixshows

	n.title = "Game of Thrones"

	n.rating = 100

	fmt.Printf("%+v/n", n)

	lion, tiger := panda()
	fmt.Println(lion, tiger)

	stringint("Loz", 420)

	var b bool
	b = true
	fmt.Println(b)

	var c bool
	fmt.Println(c)

	boobs := make(map[string]int)
	boobs["Games of Thrones"] = 1
	boobs["Suits"] = 2

	for tv, rating := range boobs {
		log.Printf("tv show %s has rating %d", tv, rating)
	}

	food := make(map[string]int)
	food["rice"] = 1
	food["beans"] = 2
	food["apple"] = 3
	delete(food, "rice")
	fmt.Println(food)

	//for loop

	d := 0
	for d < 4 {
		d++
		fmt.Println("my mum has eaten", d, "cupcakes")
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range numbers {
		fmt.Println("I love", n, "pengiuns")
		fmt.Println("I love", i, "cockatoos")
	}

	insect := make([]string, 4)
	insect[0] = "bees"
	insect[1] = "butterfly"
	insect[2] = "caterpiler"
	insect[3] = "spider"
	insectfood := make([]string, 1)
	insectfood[0] = "poo"
	copy(insectfood, insect)
	fmt.Println(insectfood)
	insect = append(insect, "lady beetle", "cockroach")
	fmt.Println(len(insect))
	insect = append(insect[:3], insect[3+1:]...)
	fmt.Println(len(insect))
	fmt.Println(insect)

	eatcuddlepandas := []int{2, 4, 6, 8}
	for i, j := range eatcuddlepandas {
		fmt.Println("I love to eat", i, "bamboos")
		fmt.Println("I love to cuddle", j, "pandas")

	}

}

func panda() (x, y string) {
	x = "Adrian is very cute"
	y = "Lauren is awesome"
	return
}

func stringint(a string, b int) {
	fmt.Printf("Loz %s has %d boobs", a, b)
}
