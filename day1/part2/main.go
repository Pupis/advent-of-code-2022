package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var result []int
	data, err := os.ReadFile("./input/day1.txt")
	if err != nil {
		fmt.Print((err))
	}

	parsedData := string(data)
	temp := strings.Split(parsedData, "\n\n")
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
		result = append(result, tempCalories)
	}
	sort.Ints(result)
	length := len(result)
	fmt.Println(result[length-1] + result[length-2] + result[length-3])
}
