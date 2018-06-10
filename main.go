package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
)

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

type person struct {
	name string

	age int

	gender string

	addressdetails addressdetails
}
type addressdetails struct {
	number string

	streetname string

	city string
}

type concert struct {
	time string

	name string

	venue string
}

type alcohol struct {
	Name string
	Ice  bool
}

type song struct {
	Title string

	Duration float64
}

type family struct {
	Names string

	Ages int

	location location
}

type location struct {
	Number int

	Street string
}

type sphere struct {
	Radius float64
}

type triangle struct {
	Base   float64
	Height float64
}

type westworld struct {
	Character string

	Rating float64
}

func (tri *triangle) changeBase() float64 {
	return 0.5 * (tri.Base * tri.Height)
}

func (sph *sphere) Volume() float64 {
	return float64(4) * math.Pi * (sph.Radius * sph.Radius)
}

func (sph *sphere) Areavolume() float64 {
	thecubed := sph.Radius * sph.Radius * sph.Radius
	return (float64(4) / float64(2)) * math.Pi * thecubed
}

func (wes *westworld) overallRating() string {
	robots := strconv.FormatFloat(wes.Rating, 'f', 1, 64)
	return wes.Character + "," + robots
}

func (m *song) summary() string {

	j := strconv.FormatFloat(m.Duration, 'f', 1, 64)
	return m.Title + ", " + j

}

type car interface {
	keyon() string
	speakeron() string
}

type cars struct {
	Driver string
	Name   string
	Sound  bool
	Broken bool
}

func (cs *cars) keyon() string {
	if cs.Broken {
		return "car will not start because key is not in the ignition"
	}
	return "car will start because key is in the ignition"
}

func (cs *cars) speakeron() string {
	if cs.Sound {
		return "speaker is ON because the keys are in the ignition"
	}
	return "speaker is OFF because the keys are in the ignition"
}

func keyon(cr car) string {
	return cr.keyon()
}

func speakeron(cr car) string {
	return cr.speakeron()
}

func main() {

	car1 := cars{
		Driver: "Panda is driving the",
		Name:   "Audi",
		Sound:  true,
		Broken: false,
	}

	car2 := cars{
		Driver: "Orangtuan is driving the",
		Name:   "Lamborghini",
		Sound:  false,
		Broken: true,
	}

	log.Println(car1.Driver, car1.Name, keyon(&car1))
	log.Println(car1.Driver, car1.Name, speakeron(&car1))
	log.Println(car2.Driver, car2.Name, keyon(&car2))
	log.Println(car2.Driver, car2.Name, speakeron(&car2))

	tri := triangle{Base: 5, Height: 3}
	fmt.Println(tri.changeBase())

	sph := sphere{
		Radius: 10,
	}
	fmt.Println(sph.Volume())
	fmt.Println(sph.Areavolume())

	wes := westworld{
		Character: "James",
		Rating:    10.76,
	}
	fmt.Println("My favourite character in Westworld is", wes.Character, "and would rate him a", wes.Rating, "out of 10.76")

	u := family{
		Names: "Georgina, Craig, Steven, Bradley, Lauren",

		Ages: 70,

		location: location{
			Number: 90,
			Street: "Bald Head",
		},
	}
	fmt.Printf("%+v\n", u)

	m := song{
		Title:    "I want to make love to you",
		Duration: 3.45,
	}
	fmt.Println(m.summary())

	var z bool
	fmt.Println(reflect.TypeOf(z))

	a := alcohol{
		Name: "Rum",
		Ice:  false,
	}
	w := alcohol{
		Name: "Vodka",
		Ice:  true,
	}
	if a != w {
		fmt.Println("a has zero ice and w has ice")
	}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", w)

	j := person{
		name:   "Bruno Mars",
		age:    28,
		gender: "male",

		addressdetails: addressdetails{
			number:     "987B",
			streetname: "Hunter Street",
			city:       "Shallowville, Sydney",
		},
	}
	fmt.Printf("%+v\n", j)

	h := boatrace{
		name:   "Fierce Sailing",
		number: 5,
		race:   "Rapid Race",
		winner: winner{
			place: 1,
			time:  0.987,
		},
	}
	fmt.Printf("%+v\n", h)

	var s dcactionheros

	s.name = "Superman"
	s.age = 34
	s.colour = "red and white"
	s.address = "4 Superman Avenue, Superville"
	fmt.Printf("%+v\n", s)

	fmt.Printf("%v\n", performance("8:00pm"))

	var pool string
	fmt.Println(reflect.TypeOf(pool))

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

	shoes := make([]string, 4)
	shoes[0] = "thongs"
	shoes[1] = "sandals"
	shoes[2] = "sandshoes"
	shoes[3] = "high heels"
	fmt.Printf("my favourite shoes are %v\n", shoes)

	q := false
	if q {
		fmt.Println("q is true")
	} else {
		fmt.Println("q is false")
	}
	y := 100
	if y >= 200 {
		fmt.Println("y is more than 200")
	} else if y <= 200 {
		fmt.Println("y is less than 150")
	}

	roads := []int{1, 2, 3, 4, 5}
	for i, r := range roads {
		fmt.Println("I have", i, r, "roads")
	}
}

func stringint(a string, b int) {
	fmt.Printf("%s has %d boobs\n", a, b)
}

func performance(time string) concert {
	f := concert{
		time:  time,
		name:  "Bruno Mars",
		venue: "Allianz Stadium",
	}
	return f
}
