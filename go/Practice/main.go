package main

type units interface {
	Celcius() float64
	Fahrenheit() float64
	Kelvin() float64
}

type kelvin struct {
	temp float64
}

type celcius struct {
	temp float64
}

type fahrenheit struct {
	temp float64
}

type myUnits struct {
	cels float64
	kelv float64
	fahr float64
}

func (k kelvin) Celcius() float64 {
	return 1234
}

func main() {

}
