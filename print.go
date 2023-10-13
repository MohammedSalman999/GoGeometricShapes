package main

import "fmt"

func PrintValue(a area) {
	area, err := a.getValue()
	if err != nil {
		fmt.Printf("Error : %v\n ",err)
		return
	}
	fmt.Printf("Area Of The Desired is  %.2f\n",area)
}