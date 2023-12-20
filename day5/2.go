package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// TODO: Loop based on source range not on a map

func Second() {
	answer := 0
	scanner := bufio.NewScanner(os.Stdin)

	var seedsRanges [][]int
	// seeds
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		seedsString := strings.Split(strings.Split(line, ":")[1], " ")[1:]
		for i := 0; i < len(seedsString)/2; i++ {
			seedRangeStartInt, _ := strconv.Atoi(seedsString[i*2])
			seedRangeCountInt, _ := strconv.Atoi(seedsString[i*2+1])
			seedsRanges = append(seedsRanges, []int{seedRangeStartInt, seedRangeCountInt})
		}
	}

	// seed to soil
	var soilsRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "seed-to-soil map:" {
			continue
		}
		destSourceCount := strings.Split(line, " ")
		dest, _ := strconv.Atoi(destSourceCount[0])
		source, _ := strconv.Atoi(destSourceCount[1])
		count, _ := strconv.Atoi(destSourceCount[2])

		var nextSeedsRanges [][]int
		for _, seedsRange := range seedsRanges {
			if source <= seedsRange[0] && seedsRange[0]+seedsRange[1] <= source+count {
				soilsRanges = append(soilsRanges, []int{dest + (seedsRange[0] - source), seedsRange[1]})
				continue
			}
			if seedsRange[0] <= source && source+count <= seedsRange[0]+seedsRange[1] {
				soilsRanges = append(soilsRanges, []int{dest, count})
				if source-seedsRange[0] > 0 {
					nextSeedsRanges = append(nextSeedsRanges, []int{seedsRange[0], source - seedsRange[0]})
				}
				if seedsRange[0]+seedsRange[1]-(source+count) > 0 {
					nextSeedsRanges = append(nextSeedsRanges, []int{source + count, seedsRange[0] + seedsRange[1] - source - count})
				}
				continue
			}
			if seedsRange[0] < source && source < seedsRange[0]+seedsRange[1] {
				soilsRanges = append(soilsRanges, []int{dest, seedsRange[0] + seedsRange[1] - source})
				nextSeedsRanges = append(nextSeedsRanges, []int{seedsRange[0], source - seedsRange[0]})
				continue
			}
			if seedsRange[0] < source+count && source+count < seedsRange[0]+seedsRange[1] {
				soilsRanges = append(soilsRanges, []int{dest + (seedsRange[0] - source), source + count - seedsRange[0]})
				nextSeedsRanges = append(nextSeedsRanges, []int{source + count, seedsRange[0] + seedsRange[1] - source - count})
				continue
			}
			nextSeedsRanges = append(nextSeedsRanges, seedsRange)
		}
		seedsRanges = nextSeedsRanges

		println("s")
		for _, s := range seedsRanges {
			println(s[0], s[1])
		}
	}
	for _, seedRanges := range seedsRanges {
		soilsRanges = append(soilsRanges, seedRanges)
	}
	// for _, s := range soilsRanges {
	// 	println(s[0], s[1])
	// }

	// soil to fertilizer
	var fertilizersRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "soil-to-fertilizer map:" {
			continue
		}
		destSourceCount := strings.Split(line, " ")
		dest, _ := strconv.Atoi(destSourceCount[0])
		source, _ := strconv.Atoi(destSourceCount[1])
		count, _ := strconv.Atoi(destSourceCount[2])

		var nextsoilsRanges [][]int
		for _, soilsRange := range soilsRanges {
			if source <= soilsRange[0] && soilsRange[0]+soilsRange[1] <= source+count {
				fertilizersRanges = append(fertilizersRanges, []int{dest + (soilsRange[0] - source), soilsRange[1]})
				continue
			}
			if soilsRange[0] <= source && source+count <= soilsRange[0]+soilsRange[1] {
				fertilizersRanges = append(fertilizersRanges, []int{dest, count})
				if source-soilsRange[0] > 0 {
					nextsoilsRanges = append(nextsoilsRanges, []int{soilsRange[0], source - soilsRange[0]})
				}
				if soilsRange[0]+soilsRange[1]-(source+count) > 0 {
					nextsoilsRanges = append(nextsoilsRanges, []int{source + count, soilsRange[0] + soilsRange[1] - source - count})
				}
				continue
			}
			if soilsRange[0] < source && source < soilsRange[0]+soilsRange[1] {
				fertilizersRanges = append(fertilizersRanges, []int{dest, soilsRange[0] + soilsRange[1] - source})
				nextsoilsRanges = append(nextsoilsRanges, []int{soilsRange[0], source - soilsRange[0]})
				continue
			}
			if soilsRange[0] < source+count && source+count < soilsRange[0]+soilsRange[1] {
				fertilizersRanges = append(fertilizersRanges, []int{dest + (soilsRange[0] - source), source + count - soilsRange[0]})
				nextsoilsRanges = append(nextsoilsRanges, []int{source + count, soilsRange[0] + soilsRange[1] - source - count})
				continue
			}
			nextsoilsRanges = append(nextsoilsRanges, soilsRange)
		}
		soilsRanges = nextsoilsRanges
	}
	for _, soilRanges := range soilsRanges {
		fertilizersRanges = append(fertilizersRanges, soilRanges)
	}

	// fertilizer to water
	var watersRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "fertilizer-to-water map:" {
			continue
		}
		destSourceCount := strings.Split(line, " ")
		dest, _ := strconv.Atoi(destSourceCount[0])
		source, _ := strconv.Atoi(destSourceCount[1])
		count, _ := strconv.Atoi(destSourceCount[2])

		var nextfertilizersRanges [][]int
		for _, fertilizersRange := range fertilizersRanges {
			if source <= fertilizersRange[0] && fertilizersRange[0]+fertilizersRange[1] <= source+count {
				watersRanges = append(watersRanges, []int{dest + (fertilizersRange[0] - source), fertilizersRange[1]})
				continue
			}
			if fertilizersRange[0] <= source && source+count <= fertilizersRange[0]+fertilizersRange[1] {
				watersRanges = append(watersRanges, []int{dest, count})
				if source-fertilizersRange[0] > 0 {
					nextfertilizersRanges = append(nextfertilizersRanges, []int{fertilizersRange[0], source - fertilizersRange[0]})
				}
				if fertilizersRange[0]+fertilizersRange[1]-(source+count) > 0 {
					nextfertilizersRanges = append(nextfertilizersRanges, []int{source + count, fertilizersRange[0] + fertilizersRange[1] - source - count})
				}
				continue
			}
			if fertilizersRange[0] < source && source < fertilizersRange[0]+fertilizersRange[1] {
				watersRanges = append(watersRanges, []int{dest, fertilizersRange[0] + fertilizersRange[1] - source})
				nextfertilizersRanges = append(nextfertilizersRanges, []int{fertilizersRange[0], source - fertilizersRange[0]})
				continue
			}
			if fertilizersRange[0] < source+count && source+count < fertilizersRange[0]+fertilizersRange[1] {
				watersRanges = append(watersRanges, []int{dest + (fertilizersRange[0] - source), source + count - fertilizersRange[0]})
				nextfertilizersRanges = append(nextfertilizersRanges, []int{source + count, fertilizersRange[0] + fertilizersRange[1] - source - count})
				continue
			}
			nextfertilizersRanges = append(nextfertilizersRanges, fertilizersRange)
		}
		fertilizersRanges = nextfertilizersRanges
	}
	for _, fertilizerRanges := range fertilizersRanges {
		watersRanges = append(watersRanges, fertilizerRanges)
	}

	// water to light
	var lightsRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "water-to-light map:" {
			continue
		}
		destSourceCount := strings.Split(line, " ")
		dest, _ := strconv.Atoi(destSourceCount[0])
		source, _ := strconv.Atoi(destSourceCount[1])
		count, _ := strconv.Atoi(destSourceCount[2])

		var nextwatersRanges [][]int
		for _, watersRange := range watersRanges {
			if source <= watersRange[0] && watersRange[0]+watersRange[1] <= source+count {
				lightsRanges = append(lightsRanges, []int{dest + (watersRange[0] - source), watersRange[1]})
				continue
			}
			if watersRange[0] <= source && source+count <= watersRange[0]+watersRange[1] {
				lightsRanges = append(lightsRanges, []int{dest, count})
				if source-watersRange[0] > 0 {
					nextwatersRanges = append(nextwatersRanges, []int{watersRange[0], source - watersRange[0]})
				}
				if watersRange[0]+watersRange[1]-(source+count) > 0 {
					nextwatersRanges = append(nextwatersRanges, []int{source + count, watersRange[0] + watersRange[1] - source - count})
				}
				continue
			}
			if watersRange[0] < source && source < watersRange[0]+watersRange[1] {
				lightsRanges = append(lightsRanges, []int{dest, watersRange[0] + watersRange[1] - source})
				nextwatersRanges = append(nextwatersRanges, []int{watersRange[0], source - watersRange[0]})
				continue
			}
			if watersRange[0] < source+count && source+count < watersRange[0]+watersRange[1] {
				lightsRanges = append(lightsRanges, []int{dest + (watersRange[0] - source), source + count - watersRange[0]})
				nextwatersRanges = append(nextwatersRanges, []int{source + count, watersRange[0] + watersRange[1] - source - count})
				continue
			}
			nextwatersRanges = append(nextwatersRanges, watersRange)
		}
		watersRanges = nextwatersRanges
	}
	for _, waterRanges := range watersRanges {
		lightsRanges = append(lightsRanges, waterRanges)
	}

	// light to temperature
	var temperaturesRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "light-to-temperature map:" {
			continue
		}
		destSourceCount := strings.Split(line, " ")
		dest, _ := strconv.Atoi(destSourceCount[0])
		source, _ := strconv.Atoi(destSourceCount[1])
		count, _ := strconv.Atoi(destSourceCount[2])

		var nextlightsRanges [][]int
		for _, lightsRange := range lightsRanges {
			if source <= lightsRange[0] && lightsRange[0]+lightsRange[1] <= source+count {
				temperaturesRanges = append(temperaturesRanges, []int{dest + (lightsRange[0] - source), lightsRange[1]})
				continue
			}
			if lightsRange[0] <= source && source+count <= lightsRange[0]+lightsRange[1] {
				temperaturesRanges = append(temperaturesRanges, []int{dest, count})
				if source-lightsRange[0] > 0 {
					nextlightsRanges = append(nextlightsRanges, []int{lightsRange[0], source - lightsRange[0]})
				}
				if lightsRange[0]+lightsRange[1]-(source+count) > 0 {
					nextlightsRanges = append(nextlightsRanges, []int{source + count, lightsRange[0] + lightsRange[1] - source - count})
				}
				continue
			}
			if lightsRange[0] < source && source < lightsRange[0]+lightsRange[1] {
				temperaturesRanges = append(temperaturesRanges, []int{dest, lightsRange[0] + lightsRange[1] - source})
				nextlightsRanges = append(nextlightsRanges, []int{lightsRange[0], source - lightsRange[0]})
				continue
			}
			if lightsRange[0] < source+count && source+count < lightsRange[0]+lightsRange[1] {
				temperaturesRanges = append(temperaturesRanges, []int{dest + (lightsRange[0] - source), source + count - lightsRange[0]})
				nextlightsRanges = append(nextlightsRanges, []int{source + count, lightsRange[0] + lightsRange[1] - source - count})
				continue
			}
			nextlightsRanges = append(nextlightsRanges, lightsRange)
		}
		lightsRanges = nextlightsRanges
	}
	for _, lightRanges := range lightsRanges {
		temperaturesRanges = append(fertilizersRanges, lightRanges)
	}

	// temperature to humidity
	var humiditysRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "temperature-to-humidity map:" {
			continue
		}
		destSourceCount := strings.Split(line, " ")
		dest, _ := strconv.Atoi(destSourceCount[0])
		source, _ := strconv.Atoi(destSourceCount[1])
		count, _ := strconv.Atoi(destSourceCount[2])

		var nexttemperaturesRanges [][]int
		for _, temperaturesRange := range temperaturesRanges {
			if source <= temperaturesRange[0] && temperaturesRange[0]+temperaturesRange[1] <= source+count {
				humiditysRanges = append(humiditysRanges, []int{dest + (temperaturesRange[0] - source), temperaturesRange[1]})
				continue
			}
			if temperaturesRange[0] <= source && source+count <= temperaturesRange[0]+temperaturesRange[1] {
				humiditysRanges = append(humiditysRanges, []int{dest, count})
				if source-temperaturesRange[0] > 0 {
					nexttemperaturesRanges = append(nexttemperaturesRanges, []int{temperaturesRange[0], source - temperaturesRange[0]})
				}
				if temperaturesRange[0]+temperaturesRange[1]-(source+count) > 0 {
					nexttemperaturesRanges = append(nexttemperaturesRanges, []int{source + count, temperaturesRange[0] + temperaturesRange[1] - source - count})
				}
				continue
			}
			if temperaturesRange[0] < source && source < temperaturesRange[0]+temperaturesRange[1] {
				humiditysRanges = append(humiditysRanges, []int{dest, temperaturesRange[0] + temperaturesRange[1] - source})
				nexttemperaturesRanges = append(nexttemperaturesRanges, []int{temperaturesRange[0], source - temperaturesRange[0]})
				continue
			}
			if temperaturesRange[0] < source+count && source+count < temperaturesRange[0]+temperaturesRange[1] {
				humiditysRanges = append(humiditysRanges, []int{dest + (temperaturesRange[0] - source), source + count - temperaturesRange[0]})
				nexttemperaturesRanges = append(nexttemperaturesRanges, []int{source + count, temperaturesRange[0] + temperaturesRange[1] - source - count})
				continue
			}
			nexttemperaturesRanges = append(nexttemperaturesRanges, temperaturesRange)
		}
		temperaturesRanges = nexttemperaturesRanges
	}
	for _, temperatureRanges := range temperaturesRanges {
		humiditysRanges = append(humiditysRanges, temperatureRanges)
	}

	// humidity to location
	var locationsRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "humidity-to-location map:" {
			continue
		}
		destSourceCount := strings.Split(line, " ")
		dest, _ := strconv.Atoi(destSourceCount[0])
		source, _ := strconv.Atoi(destSourceCount[1])
		count, _ := strconv.Atoi(destSourceCount[2])

		var nexthumiditysRanges [][]int
		for _, humiditysRange := range humiditysRanges {
			if source <= humiditysRange[0] && humiditysRange[0]+humiditysRange[1] <= source+count {
				locationsRanges = append(locationsRanges, []int{dest + (humiditysRange[0] - source), humiditysRange[1]})
				continue
			}
			if humiditysRange[0] <= source && source+count <= humiditysRange[0]+humiditysRange[1] {
				locationsRanges = append(locationsRanges, []int{dest, count})
				if source-humiditysRange[0] > 0 {
					nexthumiditysRanges = append(nexthumiditysRanges, []int{humiditysRange[0], source - humiditysRange[0]})
				}
				if humiditysRange[0]+humiditysRange[1]-(source+count) > 0 {
					nexthumiditysRanges = append(nexthumiditysRanges, []int{source + count, humiditysRange[0] + humiditysRange[1] - source - count})
				}
				continue
			}
			if humiditysRange[0] < source && source < humiditysRange[0]+humiditysRange[1] {
				locationsRanges = append(locationsRanges, []int{dest, humiditysRange[0] + humiditysRange[1] - source})
				nexthumiditysRanges = append(nexthumiditysRanges, []int{humiditysRange[0], source - humiditysRange[0]})
				continue
			}
			if humiditysRange[0] < source+count && source+count < humiditysRange[0]+humiditysRange[1] {
				locationsRanges = append(locationsRanges, []int{dest + (humiditysRange[0] - source), source + count - humiditysRange[0]})
				nexthumiditysRanges = append(nexthumiditysRanges, []int{source + count, humiditysRange[0] + humiditysRange[1] - source - count})
				continue
			}
			nexthumiditysRanges = append(nexthumiditysRanges, humiditysRange)
		}
		humiditysRanges = nexthumiditysRanges
	}
	for _, humidityRanges := range humiditysRanges {
		locationsRanges = append(locationsRanges, humidityRanges)
	}

	answer = locationsRanges[0][0]
	for i := 1; i < len(locationsRanges); i++ {
		if answer > locationsRanges[i][0] {
			answer = locationsRanges[i][0]
		}
	}
	println(answer)
}
