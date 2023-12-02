package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Part1() {
	f, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalNums := 0

	for scanner.Scan() {
		var arr []string

		line := scanner.Text()
		fmt.Println("line ", line)

		if len(line) > 0 {

			for _, v := range line {

				if unicode.IsDigit(v) {
					n := fmt.Sprintf("%c", v)
					arr = append(arr, n)
				}
			}

			fmt.Println(arr)
			numStr := arr[0] + arr[len(arr)-1]

			i, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("num ", i)
			totalNums += i
			fmt.Println(totalNums)

		}

	}

	fmt.Printf("day 1 part 1: %d\n", totalNums)

}

var dm = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func Part2() {
	f, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalNums := 0

	for scanner.Scan() {
		var arr []string

		line := scanner.Text()

		if len(line) > 0 {

			digiWord := ""
			for _, v := range line {

				if !unicode.IsDigit(v) {
					digiWord += fmt.Sprintf("%c", v)

					if val, ok := dm[digiWord]; ok {
						digiWord = ""

						arr = append(arr, val)
					} else {
						res := partString(digiWord)

						if len(res) > 0 {
							arr = append(arr, res)
							digiWord = ""
						}
					}

				} else {
					n := fmt.Sprintf("%c", v)
					arr = append(arr, n)
				}
			}

			//string of first and last digit
			fmt.Println(arr)
			var numStr string
			if len(arr) > 1 {
				numStr = arr[0] + arr[len(arr)-1]
				i, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(i)
				totalNums += i
				fmt.Println("totalnums ", totalNums)
			}

		}

	}

	fmt.Println(totalNums)
}

func partString(str string) string {
	// fmt.Println("part str", str)
	var numStr string
	for i := 0; i < len(str); i++ {
		if val, ok := dm[str[i:]]; ok {
			return val
		}
	}

	return numStr
}
