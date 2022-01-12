package main

import "fmt"

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func (c *Coordinates) SetLatitude(latitude float64) {
	c.Latitude = latitude
}

func (c *Coordinates) SetLongitude(longitude float64) {
	c.Longitude = longitude
}

func main() {
	coordinates := Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-123.08)
	fmt.Println(coordinates)
}
