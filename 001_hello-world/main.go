package main

import "fmt"

type person struct {
	fname string
	lname string
	age int
}

func( p person ) speak() {
	fmt.Println(p.fname, `says, Good Morning`)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func( sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, Bond, James Bond`)
}

func main() {
	// := is the short declaration operator
	// DECLARE a vriable and ASSIGN a value of a certain type INT, BOOL, etc.
	n, _ := fmt.Println("Hello everyone, I am studying go and hopefully I can make a bunch of money",42, true)
	fmt.Println(n)

	var x int = 23
	fmt.Printf("%T\n", x)
	
	// INFERS TYPE
	var y = 43
	fmt.Printf("%T\n", y)
	

	xi := []int{2,4,6,8,9,0}
	fmt.Println(xi)

	m := map[string]int {
		"Todd" : 34,
		"Job" : 42,
	}
	fmt.Println(m)

	p1 := person{
		"James", 
		"Major", 
		45,
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
			32,
		},
		true,
	}
	sa1.speak()
}
