package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanWords)
	ans := 0

	for scanner.Scan() {
		str := scanner.Text()
		first, last := -1, -1
		for _, char := range str {
			if unicode.IsDigit(char) {
				if first == -1{
					first = int(char - '0')	
				}
				last = int(char - '0')
			}
		}
		first*=10
		ans += first + last
	}
	fmt.Println(ans)
}

