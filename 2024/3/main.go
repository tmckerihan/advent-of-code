package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func txtToList(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cyphers []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			cyphers = append(cyphers, line)
		}
	}
	return cyphers
}

func parseMulStrings(cypher string, mulsEnabled bool, conditionsEnabled bool) ([][]int, bool, error) {
	// Combined pattern for mul(x,y), don't(), and do()
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
	matches := re.FindAllStringSubmatch(cypher, -1)

	muls := make([][]int, 0)
	for _, match := range matches {
		if match[0] == "don't()" && conditionsEnabled {
			mulsEnabled = false
		} else if match[0] == "do()" && conditionsEnabled {
			mulsEnabled = true
		} else if len(match) == 3 && mulsEnabled {
			// Handle multiplication case
			num1 := 0
			num2 := 0
			fmt.Sscanf(match[1], "%d", &num1)
			fmt.Sscanf(match[2], "%d", &num2)
			muls = append(muls, []int{num1, num2})
		}
	}
	return muls, mulsEnabled, nil
}

func puzzle1(cyphers []string) {
	total := 0
	for _, cypher := range cyphers {
		parsedMuls, _, _ := parseMulStrings(cypher, true, false)
		// fmt.Println(parsed_muls)
		for _, muls := range parsedMuls {
			total += muls[0] * muls[1]
		}
	}
	fmt.Println(total)
}

func puzzle2(cyphers []string) {
	total := 0
	mulsEnabled := true
	for _, cypher := range cyphers {
		var err error
		var parsedMuls [][]int
		parsedMuls, mulsEnabled, err = parseMulStrings(cypher, mulsEnabled, true)
		if err != nil {
			fmt.Printf("Error parsing cypher: %v\n", err)
			continue
		}
		for _, muls := range parsedMuls {
			total += muls[0] * muls[1]
		}
	}
	fmt.Println(total)
}

func main() {
	cyphers := txtToList("input.txt")
	puzzle1(cyphers)
	puzzle2(cyphers)
}
