package main

import "fmt"

func main() {

	twoSum([]int{2, 7, 11, 15}, 9)

}

func twoSum(nums []int, target int) []int {
	var i, j int
	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				fmt.Println(i, j)
			}
		}
	}
	return []int{i, j}
}
