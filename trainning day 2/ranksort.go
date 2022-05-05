// rank sort =>>> 1 2 2 3 4 4 9 ==>
//   rank 1 2 2 4 5 5 7
package main

import "fmt"

// sort mang
func sortArr(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] >= a[j] {
				c := a[i]
				a[i] = a[j]
				a[j] = c
			}
		}
	}
	return a
}

//sort rank of arrays

func rankSort(arr []int) []int {
	slice1 := make([]int, len(arr))
	var arr1 []int
	arr1 = sortArr(arr)
	// tạo 1 mảng có giá trị từ 1 tới len(arr)

	for i := 0; i < len(arr); i++ {
		slice1[i] = i + 1
	}

	//đánh dấu mang rank dựa vào mảng giá trị
	for j := 0; j < len(arr1)-1; j++ {
		if arr1[j] == arr1[j+1] {
			slice1[j+1] = j + 1
		}
	}

	// trả về mảng rank
	return slice1
}
func main() {
	var slice1 = []int{1, 2, 4, 3, 3, 4, 5}
	fmt.Println("mang", sortArr(slice1))
	print("rank ")
	fmt.Print(rankSort(slice1))
}

/*
ý tưởng mới về rank sort
rank[i]=rank[i-1]+số phần tử của [i-1]
*/
