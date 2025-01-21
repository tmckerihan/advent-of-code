package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func txtTotList(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var number_lists [][]int
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) == 0 {
			continue
		}
		numbers_line := make([]int, len(fields))
		for i, str_num := range fields {
			num, err := strconv.Atoi(str_num)
			if err != nil {
				fmt.Println("Error converting text to integer:", err)
			}
			numbers_line[i] = num
		}
		number_lists = append(number_lists, numbers_line)
	}
	return number_lists
}

func isSorted(slice []int) bool {
	isAscending := slices.IsSortedFunc(slice, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	if isAscending {
		return true
	}

	isDescending := slices.IsSortedFunc(slice, func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}
		return 0
	})

	if isDescending {
		return true
	}
	return false
}

func isSafeDistances(number_list []int) bool {
	for i := range len(number_list) - 1 {
		distance := math.Abs(float64(number_list[i] - number_list[i+1]))
		if distance < 1 || distance > 3 {
			return false
		}
	}
	return true
}

func isSafe(number_list []int) bool {
	if !isSorted(number_list) || !isSafeDistances(number_list) {
		return false
	}
	return true

}

func puzzle1(number_lists [][]int) {
	safe_count := 0
	for _, number_list := range number_lists {
		if isSafe(number_list) {
			safe_count += 1
		}
	}
	fmt.Println(safe_count)
}

func main() {
	number_lists := txtTotList("./input.txt")
	puzzle1(number_lists)
}
