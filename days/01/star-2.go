package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFilename := os.Args[1]
	file, err := os.Open(inputFilename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer file.Close()

	groupOneList := []int{}
	groupTwoList := make(map[int]int)
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
			groupTwoList[num] = groupTwoList[num] + 1
		}

		itemCount++
	}

	runningScore := 0

	for _, locationId := range groupOneList {
		runningScore += locationId * groupTwoList[locationId]
	}

	fmt.Println("Total score: ", runningScore)
}
