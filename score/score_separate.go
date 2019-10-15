package score
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//const (
//	AssignTotal = 920
//
//	highPopularity = 94
//	lowPopularity  = 1
//)
//
//// ニューリリースとそれ以外を別クラスタで処理する場合
//func main() {
//	popularity := highPopularity
//	assignNum := getAssignOfNum()
//	fmt.Println("assignNum: ", assignNum)
//
//	for i := 1; i <= AssignTotal; i++ {
//		if popularity < lowPopularity {
//			fmt.Printf("Unexpected:  %d\n", i)
//			continue
//		}
//
//		if i%assignNum == 0 {
//			fmt.Println("popularity: ", popularity)
//			popularity--
//		}
//	}
//}
//
//func getAssignOfNum() int {
//	f := math.Ceil(float64(AssignTotal) / (highPopularity - lowPopularity + 1))
//	fmt.Println("assignNum(float): ", f)
//	return int(f)
//}
