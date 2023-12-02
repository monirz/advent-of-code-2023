package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalNums := 0
	// games := make(map[string]map[string]string)

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) > 0 {

			splitted := strings.Split(line, ":")

			splittedIDs := strings.Split(splitted[0], "Game")

			sets := strings.Split(splitted[1], ";")

			red, green, blue := 0, 0, 0
			for _, v := range sets {
				rgbs := strings.Split(v, ",")

				for _, j := range rgbs {
					rgbNums := strings.Split(j, " ")

					if strings.TrimSpace(rgbNums[2]) == "red" {
						num, err := strconv.Atoi(rgbNums[1])
						if err != nil {
							log.Fatal(err)
						}

						if red < num {
							red = num
						}
					} else if strings.TrimSpace(rgbNums[2]) == "green" {
						num, err := strconv.Atoi(rgbNums[1])
						if err != nil {
							log.Fatal(err)
						}
						if green < num {
							green = num
						}
					} else if strings.TrimSpace(rgbNums[2]) == "blue" {
						num, err := strconv.Atoi(rgbNums[1])
						if err != nil {
							log.Fatal(err)
						}
						if blue < num {
							blue = num
						}
					}

				}
			}

			if red <= 12 && green <= 13 && blue <= 14 {

				fmt.Println("ids ", splittedIDs)
				n, err := strconv.Atoi(strings.TrimSpace(splittedIDs[1]))
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("valid red green blue ", red, green, blue)

				totalNums += n
			}

			fmt.Println("red green blue ", red, green, blue)

		}

	}

	fmt.Printf("day 2 part 1: %d\n", totalNums)

}

func Part2() {
	f, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	// games := make(map[string]map[string]string)

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) > 0 {

			splitted := strings.Split(line, ":")

			// splittedIDs := strings.Split(splitted[0], "Game")

			sets := strings.Split(splitted[1], ";")

			red, green, blue := 0, 0, 0
			for _, v := range sets {
				rgbs := strings.Split(v, ",")

				for _, j := range rgbs {
					rgbNums := strings.Split(j, " ")

					if strings.TrimSpace(rgbNums[2]) == "red" {
						num, err := strconv.Atoi(rgbNums[1])
						if err != nil {
							log.Fatal(err)
						}

						if red < num {
							red = num
						}
					} else if strings.TrimSpace(rgbNums[2]) == "green" {
						num, err := strconv.Atoi(rgbNums[1])
						if err != nil {
							log.Fatal(err)
						}
						if green < num {
							green = num
						}
					} else if strings.TrimSpace(rgbNums[2]) == "blue" {
						num, err := strconv.Atoi(rgbNums[1])
						if err != nil {
							log.Fatal(err)
						}
						if blue < num {
							blue = num
						}
					}

				}
			}

			fmt.Println("red green blue ", red, green, blue)
			fmt.Println("multiply ", red*green*blue)

			multiplied := red * green * blue

			sum += multiplied

		}

	}

	fmt.Printf("day 2 part 2: %d\n", sum)

}
