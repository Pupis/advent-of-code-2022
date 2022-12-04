package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input/day1.txt")
	if err != nil {
		fmt.Print((err))
	}
	parsedData := string(data)
	temp := strings.Split(parsedData, "\n\n")
	biggest := 0
	for i := 0; i < len(temp); i++ {
		calories := strings.Split(temp[i], "\n")
		tempCalories := 0
		for j := 0; j < len(calories); j++ {
			calory, err := strconv.Atoi(calories[j])
			if err != nil {
				panic(err)
			}
			tempCalories += calory
		}
		if tempCalories > biggest {
			biggest = tempCalories
		}
	}
	fmt.Println(biggest)
}
