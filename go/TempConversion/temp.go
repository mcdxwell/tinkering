package main

type units interface {
	Celsius() float64
	Fahrenheit() float64
	Kelvin() float64
}

type celsius struct{ temp float64 }
type fahrenheit struct{ temp float64 }
type kelvin struct{ temp float64 }

func (k kelvin) Celsius() float64 {
	return k.temp - 273.15
}

func (f fahrenheit) Celsius() float64 {
	return (f.temp - 32) * 5 / 9
}

func (k kelvin) Fahrenheit() float64 {
	return (k.temp-273.15)*9/5 + 32
}

func (c celsius) Fahrenheit() float64 {
	return (c.temp * 9 / 5) + 32
}

func (c celsius) Kelvin() float64 {
	return c.temp + 273.15
}

func (f fahrenheit) Kelvin() float64 {
	return (f.temp-32)*5/9 + 273.15
}
