package main

import (
	"testing"
)

var tests = []struct {
	a, b int
}{}

func TestCelsius(t *testing.T) {

	//t.Errorf()
}

func TestFahrenheit() {
	c := celsius{0}.Fahrenheit()
}

func TestKelvin() {

}
