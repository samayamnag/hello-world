package main

import (
	"fmt"
)

const (
	Pi            = 3.14
	StatusOk      = 200
	StatusCreated = 201
)

var (
	name, location, age = "Nageswara Rao", "Neelipudi", 35
)

func main() {
	fmt.Println("Hello")
	name, age, location := "Nageswara Rao", 31, "Neelipudi"

	action := func() string {
		return "Some function"
	}
	fmt.Printf("Type: %T value: %v\n", name, name)
	fmt.Printf("Type: %T value: %v \n", age, age)
	fmt.Printf("Type: %T value: %v \n", location, location)
	fmt.Printf("Type: %T, value: %v \n", action(), action())

	fmt.Printf("Type %T, value: %v\n", Pi, Pi)
	println(Pi)

	state, country := getLocation("Nag", "Bangalore")
	fmt.Println(state)
	fmt.Println(country)

	artist := &Artist{Name: "Nag", Lang: "Go", Years: 6}
	fmt.Printf("%s experince is %d\n", artist.Name, newRelease(artist))
	fmt.Printf("%s experience is %d\n", artist.Name, artist.Years)
	fmt.Printf("%s experince is %d\n", artist.Name, newRelease(artist))

	b_ptr := &name
	fmt.Println(b_ptr)
	fmt.Println(name)
	*b_ptr = "Havisa"
	fmt.Println(name)

	a_new := new(Artist)
	fmt.Println(a_new)

	var test_arr [3]string
	test_arr[0] = "Nag"
	test_arr[1] = "Phani"
	test_arr[2] = "Havisa"

	fmt.Printf("Print array of type %T: %q \n", test_arr, test_arr)

	var test_arr2 = [3]string{"Satish", "Narendra", "Samayam"}
	fmt.Printf("Print array of type %T: %q \n", test_arr2, test_arr2)
}

func getLocation(name, city string) (state, country string) {
	switch city {
	case "Bangalore":
		state, country = "Karnataka", "India"
	default:
		state, country = "Hydearabad", "India"
	}
	return
}

type Artist struct {
	Name, Lang string
	Years int
}

func newRelease(artist *Artist) int {
	artist.Years++
	return artist.Years
}