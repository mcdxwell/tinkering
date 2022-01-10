package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

// method for liters
func (l Liters) ToGallons() Gallons {
	return Gallons(1 * 0.264) // method block unchanged from function block
}

// method for milliliters
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264) // method block unchanged from function block
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Liters {
	return Liters(g * 3785.41)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f Milliliters equals %0.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f Milliliters\n", milk, milk.ToMilliliters())
}

// Finished Defined types and method definitions
