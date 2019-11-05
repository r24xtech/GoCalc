package main

import (
	"fmt"
	"strconv"
)

var counter int
var total_weight float64

func main() {
	fmt.Println("|| Welcome to ENEM#Go ||\n")

	weights := make(map[string] float64)
	grades := make(map[string] float64)

	for {
		weights["red"] = get_weight("RED")
		weights["mat"] = get_weight("MAT")
		weights["lin"] = get_weight("LIN")
		weights["hum"] = get_weight("HUM")
		weights["nat"] = get_weight("NAT")
		if counter == 5 && total_weight == 10 {
			break
		}
		counter = 0
		total_weight = 0
		fmt.Println("The sum of the weights must be 10.")
	}
	fmt.Print("\n")

	grades["red"] = get_grade("RED")
	grades["mat"] = get_grade("MAT")
	grades["lin"] = get_grade("LIN")
	grades["hum"] = get_grade("HUM")
	grades["nat"] = get_grade("NAT")

	var total_grade float64
	total_grade += weights["red"] * grades["red"]
	total_grade += weights["mat"] * grades["mat"]
	total_grade += weights["lin"] * grades["lin"]
	total_grade += weights["hum"] * grades["hum"]
	total_grade += weights["nat"] * grades["nat"]
	total_grade = total_grade/10
	fmt.Printf("\nYour final grade is: %.2f.", total_grade)
}

func get_weight(topic string) float64 {
	for{
		var weight string
		fmt.Printf("Enter the weight for %s: ", topic)
		fmt.Scanln(&weight)
		x, error := strconv.ParseFloat(weight, 64)
		if error == nil {
			counter += 1
			total_weight += x
			return x
		}
		fmt.Println("Please enter a valid weight.")
	}
}

func get_grade(topic string) float64 {
	for {
		var grade string

		fmt.Printf("Enter grade for %s: ", topic)
		fmt.Scanln(&grade)
		x, error := strconv.ParseFloat(grade, 64)
		if error == nil {
			if x >= 0 && x <= 1000 {
				return x
			}
		}
		fmt.Println("Your grade must be a number between 0 and 1000.")
	}
}
