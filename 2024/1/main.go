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

func txtToList(path string) ([]int, []int) {
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

func calculateTotalDistance(numbers_left, numbers_right []int) int {
	// sort ascending
	sort.Ints(numbers_left)
	sort.Ints(numbers_right)

	// subtract rows from eachother
	total_distance := 0
	for i := range numbers_left {
		total_distance += int(math.Abs(float64(numbers_left[i] - numbers_right[i])))

	}
	return total_distance
}

func countOccurrences(arr []int) map[int]int {
	occurences := make(map[int]int)
	for _, num := range arr {
		occurences[num]++
	}
	return occurences
}

func calculateSimilarityScore(numbers_left, numbers_right []int) int {
	occurences := countOccurrences(numbers_right)

	total_similarity := 0
	for _, val := range numbers_left {
		total_similarity += (val * occurences[val])
	}
	return total_similarity
}

func main() {
	// read columns into lists
	numbers_left, numbers_right := txtToList("./input.txt")
	// task 1
	left_copy := make([]int, len(numbers_left))
	right_copy := make([]int, len(numbers_right))
	copy(left_copy, numbers_left)
	copy(right_copy, numbers_right)
	total_distance := calculateTotalDistance(left_copy, right_copy)
	fmt.Println("The total distance is:", total_distance)

	//task 2
	total_similarity := calculateSimilarityScore(numbers_left, numbers_right)
	fmt.Println("Total similarity:", total_similarity)
}
