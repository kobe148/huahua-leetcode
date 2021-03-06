package main

import "fmt"

func main() {

	// 描述：二维数组中的每个数组的第二个元素a，在第一个一位数组中找到不大于a的数的集合，然后与二维数组中1每个数组的第一个元素求与得出最大的那个数据
	num := []int{5, 2, 4, 6, 6, 3}
	queries := [][]int{{12, 4}, {8, 1}, {6, 3}}
	arr := handler(num, queries)
	fmt.Printf("获取的集合：%v\n", arr)
}

func handler(arr1 []int, arr2 [][]int) []int {
	arr := make([]int, 3)
	if len(arr1) <= 0 || len(arr2) <= 0 {
		return arr
	}
	// 对二维数组的每个值取出与一维数组做比较，形成每个二维的一个数据集合，然后再去对比每个的运算的大小，最后保留最大的
	intSetMap := make(map[int][]int, len(arr2))
	for _, val := range arr2 {
		intSetMap[val[1]] = getAggregate(val[1], arr1)
		fmt.Printf("%d的intSetMap为：%v\n", val[1], intSetMap[val[1]])
	}
	maxArr := []int{}
	for i := 0; i < len(arr2); i++ {
		max := 0
		value := arr2[i][1]
		intMap := intSetMap[value]
		if len(intMap) <= 0 {
			//intMap = make(map[int][]int, )
		}
		for _, val := range intMap {
			intMax := arr2[i][0] ^ val
			fmt.Printf("第%d轮的数：%d与intMap为：%d的与操作结果为：%d\n", i, arr2[i][0], val, intMax)
			if intMax > max {
				max = intMax
			}
		}
		fmt.Printf("第%d轮的max为：%d\n", i, max)
		maxArr = append(maxArr, max)

	}
	return maxArr
}

func getAggregate(num int, arr []int) []int {
	set := []int{}
	if len(arr) <= 0 {
		return set
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] <= num {
			set = append(set, arr[i])
		}
	}
	return set
}
