package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)
func solve(str string, line int) int {
	
	reg:=regexp.MustCompile(`\d+`)		
	sets:=strings.Split(str, "|")
	mp:=make(map[string]int)
	for _, set := range sets {
		cardID:=""
		if line >= 1 && line <= 9 {
			cardID = "Card   " + strconv.Itoa(line) + ":" 	
		}else if line >=10 && line<=99 {
			cardID = "Card  " + strconv.Itoa(line) + ":" 	
		}else{
			cardID = "Card " + strconv.Itoa(line) + ":" 	
		}
		set = strings.TrimSpace(strings.TrimPrefix(set, cardID))
		numbers:=reg.FindAllString(set, -1)
		for _, number:= range numbers {
			mp[number]+=1	
		}
	}

	
	cnt:=0
	for _, value := range mp{
		if value > 1 {
			cnt+=1;	
		}	
	}	
	return cnt
}


func main(){
	file, _:= os.Open("input.txt")
	
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	ans, l:=0, 1
	for scanner.Scan() {
		str:=scanner.Text()
		count:=solve(str, l)
		if count > 0 {			
			ans+=int(math.Pow(2, float64(count - 1)))
		}
		l+=1
	}
	fmt.Println(ans)
}
