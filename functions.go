package main

import (
	"errors"
	"math"
)

//To Getting the area of a square we are using that function
func (s square) getValue() (float64,error) {
	if s.side<0 {
		return 0, errors.New("The Value Can't be negative ")
	}
	return s.side * s.side,nil
}

//To Getting the area of a Circle we generally use that function

func (c circle) getValue() (float64,error){
	if c.radius<0 {
		return 0, errors.New("The Value cant be negative")
	}
	return math.Pi * c.radius * c.radius,nil
}