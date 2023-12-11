package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var seedToSoil []Range
var soilToFertilizer []Range
var fertilizerToWater []Range
var waterToLight []Range
var lightToTemperature []Range
var temperatureToHumidity []Range
var humidityToLocation []Range

type Range struct {
	source      int
	destination int
	r           int
}

func main() {
	start := time.Now()
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	parseMaps(lines)
	part1(lines)
	elapsed := time.Since(start)
	log.Printf("Time %s", elapsed)
}

func parseMaps(lines []string) {
	var currentMap *[]Range
	seedToSoil = []Range{}
	soilToFertilizer = []Range{}
	fertilizerToWater = []Range{}
	waterToLight = []Range{}
	lightToTemperature = []Range{}
	temperatureToHumidity = []Range{}
	humidityToLocation = []Range{}

	for _, line := range lines {
		if line == "seed-to-soil map:" {
			currentMap = &seedToSoil
		} else if line == "soil-to-fertilizer map:" {
			currentMap = &soilToFertilizer
		} else if line == "fertilizer-to-water map:" {
			currentMap = &fertilizerToWater
		} else if line == "water-to-light map:" {
			currentMap = &waterToLight
		} else if line == "light-to-temperature map:" {
			currentMap = &lightToTemperature
		} else if line == "temperature-to-humidity map:" {
			currentMap = &temperatureToHumidity
		} else if line == "humidity-to-location map:" {
			currentMap = &humidityToLocation
		} else if line == "" {
			continue
		} else if strings.HasPrefix(line, "seeds") {
			continue
		} else {
			split := strings.Split(line, " ")
			source, _ := strconv.Atoi(split[1])
			destination, _ := strconv.Atoi(split[0])
			r, _ := strconv.Atoi(split[2])
			*currentMap = append(*currentMap, Range{source, destination, r})
		}
	}
}

func parseSeeds(lines []string) []int {
	a, _ := strings.CutPrefix(lines[0], "seeds: ")

	b := strings.Split(a, " ")
	seeds := make([]int, len(b))
	for i, seed := range b {
		seedInt, _ := strconv.Atoi(seed)
		seeds[i] = seedInt
	}

	return seeds
}

func part1(fileLines []string) {
	seeds := parseSeeds(fileLines)
	locations := []int{}
	for _, seed := range seeds {
		location := findLocation(seed)
		locations = append(locations, location)
	}

	fmt.Println("min location", findMin(locations))
}

func findMin(s []int) int {
	minimum := s[0]
	for _, el := range s {
		if el < minimum {
			minimum = el
		}
	}
	return minimum
}

func (el Range) find(source int) (int, bool) {
	if el.source <= source && source <= el.source+el.r {
		return el.destination + (source - el.source), true
	} else {
		return source, false
	}
}

func findLocation(seed int) int {
	soil := seed
	for _, el := range seedToSoil {
		res, found := el.find(seed)
		if found {
			soil = res
			break
		}
	}

	fertilizer := soil
	for _, el := range soilToFertilizer {
		res, found := el.find(soil)
		if found {
			fertilizer = res
			break
		}
	}

	water := fertilizer
	for _, el := range fertilizerToWater {
		res, found := el.find(fertilizer)
		if found {
			water = res
			break
		}
	}

	light := water
	for _, el := range waterToLight {
		res, found := el.find(water)
		if found {
			light = res
			break
		}
	}

	temperature := light
	for _, el := range lightToTemperature {
		res, found := el.find(light)
		if found {
			temperature = res
			break
		}
	}

	humidity := temperature
	for _, el := range temperatureToHumidity {
		res, found := el.find(temperature)
		if found {
			humidity = res
			break
		}
	}

	location := humidity
	for _, el := range humidityToLocation {
		res, found := el.find(humidity)
		if found {
			location = res
			break
		}
	}
	fmt.Println("location: ", location)

	return location
}
