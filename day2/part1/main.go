package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input/day2.txt")
	if err != nil {
		fmt.Print((err))
	}
	parsedData := string(data)
	temp := strings.Split(parsedData, "\n")

	points := make(map[string]int)

	points["A X"] = 3
	points["A Y"] = 6
	points["A Z"] = 0
	points["B X"] = 0
	points["B Y"] = 3
	points["B Z"] = 6
	points["C X"] = 6
	points["C Y"] = 0
	points["C Z"] = 3

	value := make(map[string]int)
	value["X"] = 1
	value["Y"] = 2
	value["Z"] = 3

	pointSum := 0

	for i := 0; i < len(temp); i++ {
		picks := strings.Split(temp[i], " ")
		you := picks[1]

		fmt.Println(points[temp[i]])
		fmt.Println(value[you])
		fmt.Println(points[temp[i]] + value[you])
		fmt.Println("------")
		roundPoint := points[temp[i]] + value[you]
		pointSum += roundPoint
		// fmt.Println()
	}

	fmt.Println(pointSum)

	// biggest := 0
	// for i := 0; i < len(temp); i++ {
	// 	calories := strings.Split(temp[i], "\n")
	// 	tempCalories := 0
	// 	for j := 0; j < len(calories); j++ {
	// 		calory, err := strconv.Atoi(calories[j])
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		tempCalories += calory
	// 	}
	// 	if tempCalories > biggest {
	// 		biggest = tempCalories
	// 	}
	// }
	// fmt.Println(biggest)
}
