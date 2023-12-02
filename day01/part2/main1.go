package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func check(e error){
	if e != nil{
		panic(e)
	}
}
func main() {
	file, err := os.Open("/home/pedrocas/Documents/AdventOfCode/2023/day01/part2/input.txt")
	check(err)
	defer file.Close()	
	scanner:= bufio.NewScanner(file) 
	scanner.Split(bufio.ScanWords)		
	var values = map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	ans := 0
	for scanner.Scan(){
		str:= string(scanner.Text())	
		numbers:= []string{"1", "one", "2", "two", "3", "three",
		"4", "four", "5", "five", "6", "six", "7", "seven", "8", 
		"eight", "9", "nine",
		}
		mp, mapLast := make(map[int]string), make(map[int]string)
		for _, number:=range numbers{
			if strings.Contains(str, number) {
				mapLast[strings.LastIndex(str, number)] = number	
				mp[strings.Index(str, number)] = number	
			}
		}
		first, last, n1, n2:= 10000000000, -1, 0, 0
		for key, _:=range mp{
			if key < first{
				first = key
			}
		}
		for key, _:=range mapLast{
			if key > last{
				last = key
			}
		}

		if len(mp[first]) == 1 {
			n1 = int(mp[first][0] - '0')
		}else {
			n1 = int(values[mp[first]] - '0')
		}
		
		if len(mapLast[last]) == 1 {
			n2 = int(mapLast[last][0] - '0')
		}else {
			n2 = int(values[mapLast[last]] - '0')
		}
		fmt.Printf("%d %d\n", n1, n2);
		n1*=10
		n1+=n2
		ans+=n1
	}
	fmt.Println(ans)
}
