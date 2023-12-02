package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cubeColor struct {
	color string
	number int
}

var possibleColors = []cubeColor {
	 {
		color: "red",
		number: 12,
	},
	 {
		color: "green",
		number: 13,
	},
	 {
		color: "blue",
		number: 14,
	},

}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Result Part 1: %v \n", partOne(lines))
	fmt.Printf("Result Part 2: %v \n", partTwo(lines))
}

func partOne(lines []string) int {
	var sum int = 0

	for _, line := range lines {
		id := getID(line)
		var gamePossible bool = true

		gamesString := line[strings.LastIndex(line, ":")+1:]

		gamesArray := strings.Split(gamesString, ";")

		for _, game := range gamesArray {
			colorsArray := strings.Split(game, ",")

			for _, color := range colorsArray {
				colorValue := strings.Split(color, " ")

				for _, colorConst := range possibleColors {
					if colorConst.color == colorValue[2] {

						colorQuantitiy, err := strconv.ParseInt(colorValue[1], 0, 0)
						if err != nil {
							log.Fatal("Error!")
						}
						if colorConst.number < int(colorQuantitiy) {
							gamePossible = false
						}
					}
				}
			}
		}

		if gamePossible {
			sum = sum + id
		}
	}

	return sum
}

func getID(line string) int {
	re := regexp.MustCompile("[0-9]+")
	idString := re.FindAllString(strings.Split(line, ":")[0], 1)[0]
	
	id, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		fmt.Printf("Error in getID")
	}
	return int(id)
}

func partTwo (lines []string) int64 {
	var sum int64 = 0

	for _, line := range lines {
		var maxRed, maxBlue, maxGreen int64 = 0, 0, 0

		gamesString := line[strings.LastIndex(line, ":")+1:]

		gamesArray := strings.Split(gamesString, ";")

		for _, game := range gamesArray {
			colorsArray := strings.Split(game, ",")

			for _, color := range colorsArray {
				colorValue := strings.Split(color, " ")

				for _, colorConst := range possibleColors {
					if colorConst.color == colorValue[2] {

						colorQuantity, err := strconv.ParseInt(colorValue[1], 0, 0)
						if err != nil {
							log.Fatal("Error!")
						}
						switch colorValue[2] {
						case "red": {
							if colorQuantity > maxRed {
								maxRed = colorQuantity
							}
						}
						case "green": {
							if colorQuantity > maxGreen {
								maxGreen = colorQuantity
							}
						}
						case "blue": {
							if colorQuantity > maxBlue {
								maxBlue = colorQuantity
							}
						}
						}
					}
				}
			}
		}

		sum = sum + (maxRed * maxGreen * maxBlue)
	}

	return sum
}