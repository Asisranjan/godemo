package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordcount := make (map[string]int)
	warray := strings.Fields(s)
	
	for _, v := range(warray){
		wordcount[v] = wordcount[v] + 1 
	}
	
	return wordcount
}

func main() {
	wc.Test(WordCount)
}
