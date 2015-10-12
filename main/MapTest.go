//package main
//
//import (
//	"golang.org/x/tour/wc"
//	"strings"
//)
//
//func WordCount(s string) map[string]int {
//	ret := make(map[string]int)
//	for _, word := range(strings.Fields(s)) {
//		value, ok := ret[word];
//		if (ok) {
//			ret[word] = value + 1
//		} else {
//			ret[word] = 1
//		}
//	}
//	return ret
//}
//
//func main() {
//	wc.Test(WordCount)
//}
