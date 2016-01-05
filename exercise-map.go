//https://tour.go-zh.org/moretypes/20
package main

import (
	"golang.org/x/tour/wc"
	"sort"
	"strings"
)

var m map[string]int

func WordCount(s string) map[string]int {
	m = make(map[string]int)
	count := 1
	res_strings := strings.Fields(s)
	sort.Strings(res_strings)
	for i := 0; i < len(res_strings); i++ {
		if i == 0 {
			m[res_strings[0]] = 1
		}
		if i >= 1 {
			if res_strings[i] == res_strings[i-1] {
				count += 1
				m[res_strings[i]] = count
				continue
			} else {
				count = 1
				m[res_strings[i]] = count
			}
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
