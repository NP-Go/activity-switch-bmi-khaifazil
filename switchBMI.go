package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert Code Here
	var weight string
	var height string

	fmt.Println("Enter your weight (kg): ")
	fmt.Scanln(&weight)
	fmt.Println("Enter your height (m): ")
	fmt.Scanln(&height)

	weightValue, _ := strconv.ParseFloat(weight, 64)
	heightValue, _ := strconv.ParseFloat(height, 64)

	switch BMI := weightValue / (heightValue * heightValue); {
	case BMI < 18.5:
		fmt.Printf("Your BMI is: %0.2f \nYou are underweight", BMI)
	case BMI >= 18.5 && BMI <= 24.9:
		fmt.Printf("Your BMI is: %0.2f \nYou are Healthy Weight", BMI)
	case BMI >= 25 && BMI <= 29.9:
		fmt.Printf("Your BMI is: %0.2f \nYou are OverWeight", BMI)
	case BMI >= 30 && BMI <= 34.9:
		fmt.Printf("Your BMI is: %0.2f \nYou are Obese", BMI)
	case BMI >= 35 && BMI <= 39.9:
		fmt.Printf("Your BMI is: %0.2f \nYou are Serverely Obese", BMI)
	default:
		fmt.Printf("Your BMI is: %0.2f \nYou are Morbidly Obese", BMI)
	}
}
