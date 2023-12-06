package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var partsTimes []string
	var partsRecords []string
	var possibleGames []int

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	partsTimes = append(partsTimes, strings.Split(lines[0], "\n")...)
	partsRecords = append(partsRecords, strings.Split(lines[1], "\n")...)

	times := strings.Fields(partsTimes[0])[1:]
	records := strings.Fields(partsRecords[0])[1:]

	for i, time := range times {
		tempsum := 0
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(records[i])
		for j := 0; j < t; j++ {
			speed := j
			distanceTraveled := speed * (t - speed)
			if distanceTraveled > d {
				tempsum += 1
			}
		}

		possibleGames = append(possibleGames, tempsum)
	}

	part1Sum := 1

	for _, game := range possibleGames {
		part1Sum *= game
	}

	timeString := ""
	for _, time := range times {
		timeString += time
	}

	recordString := ""
	for _, record := range records {
		recordString += record
	}

	timeInt, _ := strconv.Atoi(timeString)
	recordInt, _ := strconv.Atoi(recordString)

	part2Sum := 0
	for j := 0; j < timeInt; j++ {
		speed := j
		distanceTraveled := speed * (timeInt - speed)
		if distanceTraveled > recordInt {
			part2Sum += 1
		}
	}

	fmt.Println("Solution Part 1: ", part1Sum)
	fmt.Println("Solution Part 2: ", part2Sum)
}
