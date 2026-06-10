package main

import "fmt"

type PhoneNumber struct {
	CountryCode int8
	Number string
}

type Person struct {
	Name string
	Age int16
	PhoneNumbers []PhoneNumber // nested struct
}

// TODO: Use pointer for a setter function / method?
func (p *Person) SetName(name string) {
	p.Name = name
}

func (p Person) Introduce() {
	fmt.Printf("My name is %s. I'm %d.\n", p.Name, p.Age)

	fmt.Println("My phone numbers are.")

	for i := 0; i < len(p.PhoneNumbers); i++ {
		fmt.Printf("  +%d%s\n", p.PhoneNumbers[i].CountryCode, p.PhoneNumbers[i].Number)
	}
}

// Struct Composition
type Pixel struct {
	P Position
	C Color
}

type Position struct {
	X int
	Y int
}

type Color struct {
	Hex string
}

func (p Pixel) Paint() {
	fmt.Printf("The pixel at (%d, %d) is painted %s.\n", p.P.X, p.P.Y, p.C.Hex)
}

func main() {
	// var p Person = Person{
	// 	Name: "Quynh",
	// 	Age: 31,
	// 	PhoneNumbers: []PhoneNumber{
	// 		PhoneNumber{CountryCode: 84, Number: "123123123"},
	// 		PhoneNumber{CountryCode: 84, Number: "321321321"},
	// 	},
	// }

	// p.Introduce()

	// p.SetName("Oliver")

	// p.Introduce()

	var p Pixel = Pixel{
		P: Position{X: 123, Y: 456},
		C: Color{Hex: "#f0f0f0"},
	}

	p.Paint()
}
