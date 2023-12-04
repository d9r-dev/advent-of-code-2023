package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	numbers        []int64
	winningNumebrs int
	amount         int
}

type Game struct {
	numbers []int64
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

	var numbers []string

	for _, line := range lines {
		_, reading, _ := strings.Cut(line, ": ")
		numbers = append(numbers, reading)
		for _, line := range numbers {
			log.Println(line)
		}
	}

	var cards []Card
	var games []Game

	for _, cardNumbers := range numbers {
		var card Card
		var game Game
		var numbersAsInt []int64
		var winningNumbersAsInt []int64

		card.amount = 1
		scratchNumbers, winningNumbers, _ := strings.Cut(cardNumbers, " | ")

		scratchNumbersAsStrings := strings.Split(scratchNumbers, " ")

		for _, n := range scratchNumbersAsStrings {
			if n != "" {
				numberAsInt, err := strconv.ParseInt(n, 0, 0)
				if err != nil {
					log.Fatal("cannot convert string")
				}
				numbersAsInt = append(numbersAsInt, numberAsInt)
			}
		}
		card.numbers = append(card.numbers, numbersAsInt...)
		cards = append(cards, card)

		winningNumbersAsStrings := strings.Split(winningNumbers, " ")
		for _, n := range winningNumbersAsStrings {
			if n != "" {
				winningNumberAsInt, err := strconv.ParseInt(n, 0, 0)
				if err != nil {
					log.Fatal("cannot convert string")
				}

				winningNumbersAsInt = append(winningNumbersAsInt, winningNumberAsInt)
			}
		}

		game.numbers = append(game.numbers, winningNumbersAsInt...)
		games = append(games, game)
	}

	sum := 0

	for i := range games {
		points := len(findIntersection(cards[i].numbers, games[i].numbers))
		if points == 1 {
			sum += 1
		} else {
			sum += int(math.Pow(2, float64(points-1)))
		}

		if points > 0 {
			for k := 0; k < cards[i].amount; k++ {
				for j := 1; j <= points; j++ {
					cards[i+j].amount += 1
				}
			}
		}
	}

	sum2 := 0

	for i := range cards {
		sum2 += cards[i].amount
	}

	fmt.Printf("Result Part 1: %v \n", sum)
	fmt.Printf("Result Part 2: %v \n", sum2)
}

func findIntersection(arr1, arr2 []int64) []int64 {
	intersection := make([]int64, 0)

	set := make(map[int64]bool)

	for _, n := range arr1 {
		set[n] = true
	}

	for _, num := range arr2 {
		if set[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection
}
