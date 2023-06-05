package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	higherCalories := 0
	calories := 0
	data, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close() // fecha o arquivo após o fim da função

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {

			if higherCalories < calories {
				higherCalories = calories
			}

			calories = 0

		} else {
			kcl, errAtoi := strconv.Atoi(line)

			if errAtoi != nil {
				log.Fatal(errAtoi)
			}

			calories += kcl
		}
	}

	fmt.Println(higherCalories)
}
