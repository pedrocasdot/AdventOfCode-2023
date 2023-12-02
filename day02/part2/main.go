package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseSets (input string, line int) int {

	sets:=strings.Split(input, ";")
	//fmt.Println(len(sets))
	var testMap =map[string]int{
		"green" : 1,
		"blue" : 1,
		"red": 1,
	}
	for _, set := range sets {
		str := "Game " + strconv.Itoa(line) + ":" 
		//fmt.Println(str)
		set = strings.TrimSpace(strings.TrimPrefix(set, str))
		pares:= strings.Split(set, ",")

		for _, pair := range pares {
			// Separe a cor e a quantidade
			parts := strings.Fields(pair)
			if len(parts) == 2 {
				color := parts[1]
				count, err := strconv.Atoi(parts[0])
				if err == nil {
					if testMap[color] < count {
						testMap[color] = count
					}
				}
			}
		}
//		fmt.Printf("%d %d %d\n", testMap["red"], testMap["green"], testMap["blue"])	
	}
	return testMap["green"] * testMap["blue"] * testMap["red"]   
}

func main(){		
	file, err:= os.Open("input.txt")
	check(err)
	
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	cnt, ans:=1, 0
	for scanner.Scan() {
		str := scanner.Text()
		value:=parseSets(str, cnt)
		ans+=value	
		cnt+=1	
	}
	fmt.Println(ans)
}

