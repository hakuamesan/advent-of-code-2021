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
	var threeWindow []int

	for n := 0; n < 1998; n++ {
		sum := result[n] + result[n+1] + result[n+2]
		threeWindow = append(threeWindow, sum)
		println("n=", threeWindow[n], n)
	}

	for n := 0; n < 1997; n++ {
		println("sum=", threeWindow[n], n)
		if threeWindow[n] < threeWindow[n+1] {
			count++
		}
	}

	println("Answer part2:", count)
}
