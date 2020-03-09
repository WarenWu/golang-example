package ex

import "fmt"

//new使用
func Ex()  {
	var arryLP = new([3]int)
	var sliceLP = new([]int)
	var mapLP = new(map[int]int)

	
	arryLP[0] = 1
	//new 不支持引用类型，用于数组和结构体
	//sliceLP[0] = 2
	//mapLP[0] = 3

	fmt.Println(arryLP, sliceLP, mapLP)
}