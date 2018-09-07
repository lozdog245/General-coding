package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
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
	fuel() int
}

type cars struct {
	Driver string
	Name   string
	Sound  bool
	Broken bool
	Fuels  int
}

func (cs *cars) fuel() int {
	if cs.Fuels <= 50 {
		return 0
	}
	if cs.Fuels >= 50 {
		return 100
	}
	return 55
}

func (cs *cars) keyon() string {
	if cs.Broken {
		return "car will not start because there is no fuel in in the ignition"
	}
	return "car will start because key is in the ignition"
}

func (cs *cars) speakeron() string {
	if cs.Sound {
		return "speaker is ON because the keys are in the ignition"
	}
	return "speaker is OFF because the keys are not in the ignition"
}

func keyon(cr car) string {
	return cr.keyon()
}

func speakeron(cr car) string {
	return cr.speakeron()
}

func fuel(cr car) int {
	return cr.fuel()
}

// ToLower changing strings to lower case
func ToLower(str string) string {
	return strings.ToLower(str)
}

//ToUpper changing strings to upper case
func ToUpper(upp string) string {
	return strings.ToUpper(upp)
}

// Replace functions
func Replace(spacey string, space string, stars string, loz int) string {
	return strings.Replace(spacey, space, stars, loz)
}

// TrimSpace - fixes spaces in strings
func TrimSpace(trim string) string {
	return strings.TrimSpace(trim)
}

// checking if a number can be halved

func half(halftohalf int) (int, error) {
	if halftohalf%2 != 0 {
		return -1, fmt.Errorf("the number %v cannot be halved", halftohalf)
	}
	return halftohalf / 2, nil
}

func slowfunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("I am still sleeping for another 2 seconds")
	// time.Sleep pauses the execution of a program for a particular duration (* 2 seconds)

}

// working with channels

func slowreaction(c chan string) {
	time.Sleep(time.Second * 5)
	c <- "slow reaction() finished"
}

func slowy(p chan string) {
	for messages := range p {
		fmt.Println(messages)
	}
}

func responsetime(urls string) {
	start := time.Now()

	website, error := http.Get(urls)

	if error != nil {
		fmt.Println(error)
	}

	defer website.Body.Close()

	elapsed := time.Since(start).Seconds()

	// time.since refers to the parameter which we specify ie - time.Now.

	fmt.Printf(" the wesbite %s took %v seconds to download\n", urls, elapsed)
}

func slowpo(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}
}

func ping1(z chan string) {
	time.Sleep(time.Second * 3)
	z <- "ping1 on chanel5"
}

func ping2(q chan string) {
	time.Sleep(time.Second * 4)
	q <- "ping2 on chanel6"
}

func sender(c chan string) {
	t := time.NewTicker(2 * time.Second)
	for {
		c <- "time is running out"
		<-t.C
	}
}

func main() {

	messaging := make(chan string)
	stop := make(chan bool)

	go sender(messaging)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println(messaging)
		stop <- true
	}()

	chanel5 := make(chan string)
	chanel6 := make(chan string)

	go ping1(chanel5)
	go ping2(chanel6)

	select {
	case msg1 := <-chanel5:
		fmt.Println("received", msg1)
	case msg2 := <-chanel6:
		fmt.Println("received", msg2)
		// add if no receivers found example here:
		// case <-time.After(500 * time.Millisecond):
		// fmt.Println("there was no responses")
	}

	message := make(chan string)
	go slowpo(message)
	messy := <-message
	fmt.Println(messy)

	//channels
	c := make(chan string)
	go slowreaction(c)

	msg := <-c // <- channel syntax
	fmt.Println(msg)

	messages := make(chan string, 2) // (contains sender and receiver)
	messages <- "hello"
	messages <- "sexy pie"
	close(messages)
	fmt.Println("there is currently no receivers")
	time.Sleep(time.Second * 4)
	for receiver := range messages { // the receiver function iterates over the channel using range and prints the buffered messages in the channel to console
		fmt.Println(receiver)
	}

	slowfunc()
	fmt.Println("I am awake now")

	urls := make([]string, 3)
	urls[0] = "https://arstechnica.com/"
	urls[1] = "https://litecoin.org/"
	urls[2] = "http://litecoinblockhalf.com/"

	for _, url := range urls {
		go responsetime(url)
	}

	time.Sleep(time.Second * 5)

	// func response time

	// blocking function - only write the function name and then print it - example here

	// using go keyword can unblock a function
	// unblocking a function example below:
	// 			go slowfunc()
	// 			fmt.Println("i am now  shown straight away")
	// 			time.Sleep(time.Second * 3) (by adding time.sleep again to the function main allows the program to continue to execute before slowfunc completes)

	mistake, numberino := half(101)
	if numberino != nil {
		fmt.Println(numberino)
	}
	fmt.Println(mistake)

	// creating errors errors.New.

	opps := errors.New("the password entered is incorrect, please try again")
	if opps != nil {
		fmt.Println(opps)
	}

	// DO NOT USE THE PANIC FUNCTION

	// an error is created using the new method. errors.New

	idcheck, number := "passport", 19
	id := fmt.Errorf("Wilson's %v (%d code error detected - passport has expired), please print new ticket", idcheck, number)
	if id != nil {
		fmt.Println(id)
	}

	file, err := ioutil.ReadFile("laugh.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", file)

	fmt.Println(TrimSpace(" HELLO BIG BAFOON      "))

	fmt.Println(ToLower("PLEASE RETURN THIS TO LOWER CASE"))

	fmt.Println(ToUpper("please return to upper case"))

	car1 := cars{
		Driver: "Panda",
		Name:   "Lamborghini",
		Broken: false,
		Sound:  true,
		Fuels:  100,
	}

	car2 := cars{
		Driver: "Orangtuan",
		Name:   "Lamborghini",
		Sound:  false,
		Broken: true,
		Fuels:  0,
	}

	log.Printf("The car's name is %s and the driver is %s. The fuel is %d%%.", car1.Name, car1.Driver, car1.Fuels)

	log.Printf("%s %s %s %d %s\n %s\n %s\n",
		car1.Driver, car1.Name, "there is", car1.Fuels, "percent of fuel in the car", keyon(&car1), speakeron(&car1))
	log.Printf("%s %s %s %d %s\n %s\n %s\n",
		car2.Driver, car2.Name, "there is", car2.Fuels, "percent of fuel in the car", keyon(&car2), speakeron(&car2))

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

	var v bool
	fmt.Println(v)

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

	i := "I am an interpreted string literal"
	fmt.Println(i)

	l := "Panda lives in Pandaville \n Panda is 15 years old \n  Panda cooks bamboo for a living \t"
	fmt.Println(l)

	bh := `I love this autumn weather 
	The leaves are falling from the trees 
		Autum is my favourtie season`
	fmt.Println(bh)

	e := "Can you hear me?"
	e += "\n Hear me screaming ' ?"
	fmt.Println(e)

	var o = 6
	var r = " doughnuts and cream "
	intToString := strconv.Itoa(o)
	var dessert string = intToString + r
	fmt.Println(dessert)

	var buffer bytes.Buffer

	for l := 10; l < 20; l++ {
		buffer.WriteString(" panda ")
	}
	fmt.Println(buffer.String())

	fmt.Println(strings.Index("surface", "face"))
	fmt.Println(strings.Index("moon", "aer"))

	fmt.Println(strings.TrimSpace(" I don't need all this space "))

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
