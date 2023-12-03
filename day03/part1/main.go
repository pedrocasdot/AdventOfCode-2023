package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	//"strings"
)

func checkValid(i int, j int, tamI int, tamJ int, table[]string) bool{
	return i>=0 && j>=0 && i < tamI && j < tamJ && (table[i][j] != '.' && !(table[i][j]>='0' && table[i][j]<='9'))	
}

func solve(s string, line int,  table[]string) int{
	dx:= [8]int{-1, 0, 1,  0, -1, 1,-1,1}
    dy := [8]int{ 0, 1, 0, -1, -1, 1,1,-1}
	numberRows, numberCol:=len(table),len(table[0]) 
	//fmt.Printf("[%d][%d]\n", numberRows, numberCol);
	re:=regexp.MustCompile(`\d+`)
	matches:=re.FindAllStringIndex(s, -1)
	matchesV:=re.FindAllString(s, -1)
	var pos [][2]int
	for _, match := range matches{
		pos = append(pos, [2]int{match[0], match[1] - 1})
	}
	row, sum, cnt:=0, 0, 0
	for _, match := range matchesV{

		valid:=false
		for i := pos[row][0] ;i <= pos[row][1]; i++{
			for d:=0; d < 8; d++ {
				if checkValid(line + dx[d], i + dy[d], numberRows, numberCol, table) {
					valid = true			
				}		
			}	
		}
		if valid {
			cnt+=1
			val, _ := strconv.Atoi(match)
			sum+=val
		}
		row+=1	
	}
	return sum
}
func main(){
	
	file,_:=os.Open("input.txt")
	
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var table []string
	for scanner.Scan() {
		str := scanner.Text() 
		if len(str) > 0 {
			table = append(table, str)
		}
	}
	ans:=0
	for i:=0 ; i < len(table); i++{
		ans+=solve(table[i], i, table)	
	}
	fmt.Println(ans)
}
