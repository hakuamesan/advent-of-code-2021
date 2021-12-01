package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	file_content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(file_content)))
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, x)
	}

	count := 0
	for n, _ := range result {
		if n < 1999 {
			println("n=", result[n], n)
			if result[n+1] > result[n] {
				count++
			}
		}
	}

	println("Answer part1:", count)
}
