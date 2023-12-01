package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
	"strings"
)

var NUMBERS = map[string]string{
	"one": "o1one",
	"two": "t2two",
	"three": "t3thre",
	"four": "f4four",
	"five": "f5five",
	"six": "s6six",
	"seven": "s7seven",
	"eight": "e8eight",
	"nine": "n9nine",
	"zero": "z0zero",
}

func main(){
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var replacedLines []string

	for _, v := range lines {
		replacedLines = append(replacedLines, replaceStringsByDigits(v))
	}

	
	fmt.Printf("Result Part1: %v \n", makeSum(lines))
	fmt.Printf("Result Part2: %v", makeSum(replacedLines))
}

func makeSum(lines []string) int {
	var sum int

	for _, s := range lines {
		
		sum = sum + int(handleOnlyDigits(findDigitsAndCreateConfigValueString(s)))
	}

	return sum
}

func findDigitsAndCreateConfigValueString(s string) string {
	var value string

	for j := 0; j < len(s); j++ {
			var char rune = rune(s[j])
			if unicode.IsDigit(char) {
				value = value + string(char)
			} 
		} 

	return value
}

func handleOnlyDigits(value string) int {
	var valueToAdd string = ""
	if (len(value) == 1) {
		valueToAdd = value + value
	} else {
		valueToAdd = string(value[0]) + string(value[len(value)-1])
	}

	valueAsInt,err := strconv.ParseInt(valueToAdd, 0, 16)
	if err != nil {
		log.Fatal(err)
	}

	return int(valueAsInt)
}

func replaceStringsByDigits(input string) string {
	val := input

	for k, r := range NUMBERS {
		val = strings.ReplaceAll(val, k, r)
	}

	return val
}