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

func parseSets (input string, line int) bool {

	sets:=strings.Split(input, ";")
	//fmt.Println(len(sets))
	veri:=true
	for _, set := range sets {
		str := "Game " + strconv.Itoa(line) + ":" 
		//fmt.Println(str)
		set = strings.TrimSpace(strings.TrimPrefix(set, str))
		pares:= strings.Split(set, ",")
		testMap:=make(map[string]int)
		for _, pair := range pares {
			// Separe a cor e a quantidade
			parts := strings.Fields(pair)
			if len(parts) == 2 {
				color := parts[1]
				count, err := strconv.Atoi(parts[0])
				if err == nil {
					testMap[color]+=count
				}
			}
		}
		veri = veri && testMap["red"] <= 12 && testMap["green"] <= 13 && testMap["blue"]<=14
//		fmt.Printf("%d %d %d\n", testMap["red"], testMap["green"], testMap["blue"])	
	}
	return veri    
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
		if value {
			ans+=cnt
		}
		cnt+=1	
	}
	fmt.Println(ans)
}
