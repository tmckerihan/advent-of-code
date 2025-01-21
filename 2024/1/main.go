package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func txt_to_list(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers_left []int
	var numbers_right []int

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "   ")
		if len(fields) == 2 {
			num, err := strconv.Atoi(fields[0])
			if err != nil {
				fmt.Println("Error converting text to integer:", err)
				continue
			}
			numbers_left = append(numbers_left, num)

			num, err = strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Error converting text to integer:", err)
				continue
			}
			numbers_right = append(numbers_right, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return numbers_left, numbers_right
}

func main() {
	// read columns into lists
	numbers_left, numbers_right := txt_to_list("./input.txt")

	// sort ascending
	sort.Ints(numbers_left)
	sort.Ints(numbers_right)

	// subtract rows from eachother
	total_distance := 0
	for i := range numbers_left {
		total_distance += int(math.Abs(float64(numbers_left[i] - numbers_right[i])))

	}

	fmt.Println("The total distance is: ", total_distance)
}
