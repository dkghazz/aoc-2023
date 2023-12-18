package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func First() {
	answer := 0
	scanner := bufio.NewScanner(os.Stdin)

	var seeds []int
	// seeds
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		seedsString := strings.Split(strings.Split(line, ":")[1], " ")[1:]
		for _, seed := range seedsString {
			seedInt, _ := strconv.Atoi(seed)
			seeds = append(seeds, seedInt)
			// println(seedInt)
		}
	}

	// seed to soil
	soils := make(map[int]int, len(seeds))
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

		for i, seed := range seeds {
			if source <= seed && seed <= source+count {
				soils[i] = dest + (seed - source)
			}
		}
	}
	for i := 0; i < len(seeds); i++ {
		if soils[i] == 0 {
			soils[i] = seeds[i]
		}
	}

	// soil to fertilizer
	fertilizers := make(map[int]int, len(seeds))
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

		for i, soil := range soils {
			if source <= soil && soil <= source+count {
				fertilizers[i] = dest + (soil - source)
			}
		}
	}
	for i := 0; i < len(soils); i++ {
		if fertilizers[i] == 0 {
			fertilizers[i] = soils[i]
		}
		println(fertilizers[i])
	}

	// fertilizer to water
	waters := make(map[int]int, len(seeds))
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

		for i, fertilizer := range fertilizers {
			if source <= fertilizer && fertilizer <= source+count {
				waters[i] = dest + (fertilizer - source)
			}
		}
	}
	for i := 0; i < len(fertilizers); i++ {
		if waters[i] == 0 {
			waters[i] = fertilizers[i]
		}
	}

	// water to light
	lights := make(map[int]int, len(seeds))
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

		for i, water := range waters {
			if source <= water && water <= source+count {
				lights[i] = dest + (water - source)
			}
		}
	}
	for i := 0; i < len(waters); i++ {
		if lights[i] == 0 {
			lights[i] = waters[i]
		}
	}

	// light to temperature
	temperatures := make(map[int]int, len(seeds))
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

		for i, light := range lights {
			if source <= light && light <= source+count {
				temperatures[i] = dest + (light - source)
			}
		}
	}
	for i := 0; i < len(lights); i++ {
		if temperatures[i] == 0 {
			temperatures[i] = lights[i]
		}
	}

	// temperature to humidity
	humiditys := make(map[int]int, len(seeds))
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

		for i, temperature := range temperatures {
			if source <= temperature && temperature <= source+count {
				humiditys[i] = dest + (temperature - source)
			}
		}
	}
	for i := 0; i < len(temperatures); i++ {
		if humiditys[i] == 0 {
			humiditys[i] = temperatures[i]
		}
	}

	// humidity to location
	locations := make(map[int]int, len(seeds))
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

		for i, humidity := range humiditys {
			if source <= humidity && humidity <= source+count {
				locations[i] = dest + (humidity - source)
			}
		}
	}
	for i := 0; i < len(humiditys); i++ {
		if locations[i] == 0 {
			locations[i] = humiditys[i]
		}
	}

	answer = locations[0]
	for i := 1; i < len(locations); i++ {
		if answer > locations[i] {
			answer = locations[i]
		}
	}
	println(answer)
}
