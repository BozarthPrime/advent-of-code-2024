package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/******************************************
 * Requirements:
 * - As each location is checked they will mark it on their list with a star
 * - They figure the chief historian _must_ be in one of the first 50 places
 * - Must help them get fifty starts before Santa takes off on Dec. 25th
 * - Two puzzles each day, each grants one star
 * - List of locations is currently empty
 * - Fist place to check is the Chief Historian's office
 *   - On checking the office all agree that the Chief Historian is no where to be found
 *   - An assortment of notes and historically significant locations were found
 *   - This was the planning the Chief historian was doing before they left
 * - Historically significant locations are listed by UUID (location ID) instead of name
 * - They have split into two groups, each of which is trying to find a complete list of all location IDs in the office
 * - The compared lists (input) do not look similar
 * - We need to reconcile the lists, presented as rows of 2 numbers, one from each group
 * - Pair numbers in each column ascending
 * - Within each pair, determine the distance (difference)
 * - Sum differences
 ******************************************/

func main() {
	inputFilename := os.Args[1]
	file, err := os.Open(inputFilename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer file.Close()

	groupOneList := []int{}
	groupTwoList := []int{}
	itemCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Could not parse int from: ", scanner.Text())
			panic(err)
		}

		if itemCount%2 == 0 {
			groupOneList = append(groupOneList, num)
		} else {
			groupTwoList = append(groupTwoList, num)
		}

		itemCount += 1
	}

	sort.Ints(groupOneList)
	sort.Ints(groupTwoList)

	runningDiff := 0

	for i := 0; i < itemCount/2; i++ {
		runningDiff += getDiff(groupOneList[i], groupTwoList[i])
	}

	fmt.Println("Total diff: ", runningDiff)
}

func getDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}
