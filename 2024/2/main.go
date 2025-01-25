package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func txtToList(path string) [][]int {
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

func isAscending(slice []int, allow_skip bool) (bool, int) {
	skipped_idx := -1
	for i := range len(slice) - 1 {
		// ascending check
		if slice[i] < slice[i+1] {
			continue
		} else if allow_skip {
			// try removing current idx
			if (i < len(slice)-2) && slice[i+1] < slice[i+2] {
				if skipped_idx != -1 {
					// skip already used
					return false, -1
				}
				skipped_idx = i
				continue
			} else if (i < len(slice)-2) && (slice[i] < slice[i+2]) {
				// try removing next idx
				if skipped_idx != -1 {
					// skip already used
					return false, -1
				}
				skipped_idx = i + 1
				continue
			} else {
				if skipped_idx != -1 {
					// skip already used
					return false, -1
				}
				skipped_idx = i + 1
				continue
			}
		} else {
			return false, -1
		}
	}
	return true, skipped_idx
}

func isDescending(slice []int, allow_skip bool) (bool, int) {
	skipped_idx := -1
	for i := range len(slice) - 1 {
		// descending check
		if slice[i] > slice[i+1] {
			continue
		} else if allow_skip {
			// try removing current idx
			if (i < len(slice)-2) && slice[i+1] > slice[i+2] {
				if skipped_idx != -1 {
					// skip already used
					return false, -1
				}
				skipped_idx = i
				continue
			} else if (i < len(slice)-2) && (slice[i] > slice[i+2]) {
				if skipped_idx != -1 {
					// skip already used
					return false, -1
				}
				skipped_idx = i + 1
				continue
			} else {
				if skipped_idx != -1 {
					// skip already used
					return false, -1
				}
				skipped_idx = i + 1
				continue
			}
		} else {
			return false, -1
		}
	}
	return true, skipped_idx
}

func isSorted(slice []int, allow_skip bool) (bool, int) {
	ascendingCheck, skipped_idx := isAscending(slice, allow_skip)
	if ascendingCheck {
		return true, skipped_idx
	}
	return isDescending(slice, allow_skip)
}

func isSafeDistances(number_list []int, allow_skip_idx int) bool {
	skipped_idx := -1
	for i := range len(number_list) - 1 {
		if i == skipped_idx || allow_skip_idx == i || allow_skip_idx == i+1 {
			// skip unnecessary idx
			continue
		}
		distance := math.Abs(float64(number_list[i] - number_list[i+1]))
		if distance < 1 || distance > 3 {
			// if skip allowed
			if (allow_skip_idx == -1 || allow_skip_idx == i+1) && i < len(number_list)-2 {
				// allow single skip
				if skipped_idx != -1 && skipped_idx != i+1 {
					return false
				}
				skipped_idx = i + 1
				distance = math.Abs(float64(number_list[i] - number_list[i+2]))
				if distance < 1 || distance > 3 {
					return false
				}
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func isSafe(number_list []int, allow_skip bool) bool {
	listIsSorted, skipped_idx := isSorted(number_list, allow_skip)
	if !listIsSorted {
		return false
	}
	// allow a single rule skip
	listIsSafeDistances := isSafeDistances(number_list, skipped_idx)
	if !listIsSafeDistances {
		fmt.Println(number_list, skipped_idx)
	}
	return listIsSafeDistances
}

func puzzle1(number_lists [][]int) {
	safe_count := 0
	for _, number_list := range number_lists {
		if isSafe(number_list, false) {
			safe_count += 1
		}
	}
	fmt.Println(safe_count)
}

func puzzle2(number_lists [][]int) {
	safe_count := 0
	for _, number_list := range number_lists {
		if isSafe(number_list, true) {
			fmt.Println(number_list, "safe")
			safe_count += 1
		}
	}
	fmt.Println(safe_count)
}

func main() {
	number_lists := txtTotList("./input.txt")
	// puzzle1(number_lists)
	puzzle2(number_lists)
}
