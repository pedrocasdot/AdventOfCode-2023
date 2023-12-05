package main

import (
	"bufio"
	"fmt"
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

	
	var cnt int
	for _, value := range mp{
		if value > 1 {
			cnt+=1;	
		}	
	}	
	return cnt
}

var ans map[int]int
func dfs(u int, adj map[int][]int){
	ans[u]+=1
	for _, v := range adj[u]{
		dfs(v, adj)
	}
}

func main(){
	file, _:= os.Open("input.txt")
	
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	l, sum:=0, 0
	
	cardsMatch, card, adj:=make(map[int]int), make(map[int]int), make(map[int][]int)
	ans = make(map[int]int)
	for scanner.Scan() {
		str:=scanner.Text()
		count:=solve(str, l + 1)
		cardsMatch[l + 1] = count
		l+=1
	}
		
	for i:=1; i<=l; i++ {card[i] = 1}
	for i:=1; i<=l; i++ {
		for j :=  1; j<=cardsMatch[i] ; j++ {
			adj[i] = append(adj[i], i + j)
		}
	}
	for i:= 1; i<=l; i++{
		dfs(i, adj)
	}
	for i := 1; i<=l; i++{
		sum+=ans[i]
	}
	fmt.Println(sum)
}
