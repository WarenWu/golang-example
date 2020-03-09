package ex

import (
	"fmt"
	"reflect"
)

//map的使用
func Ex() {

	//定义和初始化
	var map_l = make(map[int][]float32)  //map[key]value 中value可以是数组/切片/map/基本类型
	map_l[1] = []float32{0: 1.0, 1: 1.1} // 切片类型也可以使用key模式初始化，即{key:value}
	map_l[2] = []float32{0: 2.0, 1: 2.1, 2: 2.2}
	map_l[3] = []float32{0: 3.0, 1: 3.1, 2: 3.2, 3: 3.3}

	//打印值类型
	fmt.Println("reflect:", reflect.TypeOf(map_l))
	//fmt.Printf("printf:%t\n", map_l)

	//只获取key
	for key := range map_l {
		fmt.Print(key, ",")
	}
	//获取value
	for key, v := range map_l {
		fmt.Printf("\n[%d]:%#v", key, v)
	}
}
