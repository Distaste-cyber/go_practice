package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	// *p = 10 assigns 10 at the address stored in the pointer
	fmt.Printf("The value p points to is: %v \n", *p)
	fmt.Printf("The value pf i is: %v \n", i)
	p = &i //Assign the address of i to pointer p

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of thing 1 %p \n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\n The result is: %v \n", result)

}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing 2 %p \n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}
