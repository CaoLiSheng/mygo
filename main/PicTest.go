//package main
//
//import "golang.org/x/tour/pic"
//
//func Pic(dx, dy int) [][]uint8 {
//	var dots [][]uint8
//
//	for i := 0; i < 256; i++ {
//		var line []uint8
//		for j := 0; j < 256; j++ {
//			line = append(line, uint8((i + j) / 2))
//		}
//		dots = append(dots, line)
//	}
//
//	return dots
//}
//
//func main() {
//	pic.Show(Pic)
//}
