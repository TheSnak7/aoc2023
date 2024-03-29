package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() int {
	//return calculateIdSum("../day2/input.txt", 12, 13, 14)
	return calculatePower("../day2/input.txt")
}

func calculatePower(filePath string) int {

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	powerSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.TrimSpace(strings.Split(line, ":")[1])
		rounds := strings.Split(game, ";")

		var rMin, gMin, bMin int

		for _, round := range rounds {
			red, green, blue := parseRound(round)

			rMin = max(red, rMin)
			gMin = max(green, gMin)
			bMin = max(blue, bMin)

		}

		power := rMin * gMin * bMin
		powerSum += power
	}

	file.Close()

	return powerSum
}

func calculateIdSum(filePath string, rMax int, gMax int, bMax int) int {
	idSum := 0

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	gameId := 1

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.TrimSpace(strings.Split(line, ":")[1])
		rounds := strings.Split(game, ";")

		gamePossible := true

		for _, round := range rounds {
			red, green, blue := parseRound(round)

			if (red > rMax) || (green > gMax) || (blue > bMax) {
				gamePossible = false
			}
		}

		if gamePossible {
			idSum += gameId
		}

		gameId += 1
	}

	file.Close()

	return idSum
}

func parseRound(round string) (red int, green int, blue int) {
	cubes := strings.Split(round, ",")

	red = 0
	green = 0
	blue = 0

	for _, val := range cubes {
		// val looks like this: " 5 red"
		statements := strings.Split(strings.TrimSpace(val), " ")
		num, err := strconv.ParseInt(statements[0], 10, 32)

		if err != nil {
			fmt.Println(err)
		}

		switch statements[1] {
		case "red":
			{
				red = int(num)
			}
		case "green":
			{
				green = int(num)
			}
		case "blue":
			{
				blue = int(num)
			}

		}
	}

	return red, green, blue
}
