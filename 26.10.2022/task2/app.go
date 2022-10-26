package main

import "fmt"


func FindMinimal(a []int) int{
	var min int = a[0]
	mp := make(map[int]int)
	for _, val := range a{
		if min > val{
			min = val
		}
		mp[val]++
	}
	fmt.Println(mp)
	return mp[min]
}

func main(){
	a := []int{4,5,6,7,8,9,7,7,7,6,4}
	fmt.Println(FindMinimal(a))
}