package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cell struct {
	digit int
	countOfRepetitions int
}
func main() {
	const count_of_digits = 10
	counter := []cell{}
	for i:=0;i<count_of_digits;i++{
		localCell := cell{i,0}
		counter = append(counter, localCell)
	}
	var str string
	sc := bufio.NewScanner(os.Stdin)
	println("Input ints string")
	sc.Scan()
	str += sc.Text()
	for i:=0 ; i< count_of_digits; i++{
		counter[i].countOfRepetitions = strings.Count(str,strconv.Itoa(i))
	}
	sort.Slice(counter, func(i, j int) bool {
		return counter[i].countOfRepetitions >= counter[j].countOfRepetitions
	})
	for _,cell :=range(counter){
		for k:=0;k < cell.countOfRepetitions;k++{
			print(cell.digit)
		}
	}
}