package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	//"strings"
)

func checkValid(i int, j int, tamI int, tamJ int) bool{
	return i>=0 && j>=0 && i < tamI && j < tamJ;	
}

type Tuple struct {
    First  int
    Second int
	Row int
}

var mp map[Tuple]string
var pos [][3]int
var row int
func solve(s string, line int){
	re:=regexp.MustCompile(`\d+`)
	matches:=re.FindAllStringIndex(s, -1)
	matchesV:=re.FindAllString(s, -1)

	for _, match := range matches{
		pos = append(pos, [3]int{match[0], match[1] - 1, line})
	}
	for _, match := range matchesV{
		pair1:=Tuple{pos[row][0], pos[row][1], line}
		mp[pair1] = match
		row+=1	
	}
}
func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}
func main(){
	dx:= [8]int{-1, 0, 1,  0, -1, 1,-1,1}
    dy := [8]int{ 0, 1, 0, -1, -1, 1,1,-1}
	mp = make(map[Tuple]string)
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
		solve(table[i], i)	
	}
	for i:=0 ; i < len(table); i++{
		for j:= 0; j < len(table[0]); j++{
			if  table[i][j] == '*' {
				vis:= make(map[Tuple]bool)
				var numbers []int
				for d:=0; d < 8; d++ {
					nx, ny:=i+dx[d], j + dy[d]
					if checkValid(nx, ny, len(table), len(table[0])) {
						for k:= 0; k < len(pos); k++{
							if ny>=pos[k][0] && ny<=pos[k][1] && pos[k][2] == nx && vis[Tuple{pos[k][0], pos[k][1], pos[k][2]}] == false{
								value,_:=strconv.Atoi(mp[Tuple{pos[k][0], pos[k][1], pos[k][2]}])
								numbers = append(numbers, value)
								vis[Tuple{pos[k][0], pos[k][1], pos[k][2]}] = true
							}		
						}		
					}		
				}
				if  len(numbers) > 1{
					for d:=0; d <len(numbers); d++{
						for e := d + 1; e < len(numbers); e++{
							ans+=numbers[d] * numbers[e]
						}
					}			
				}
			}	
		}
	}
	fmt.Println(ans)
}

