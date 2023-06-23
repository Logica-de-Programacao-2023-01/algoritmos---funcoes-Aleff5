package main

import (
	"fmt"
	"sort"
)

func findSecondLargest(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, fmt.Errorf("slice deve conter pelo menos dois elementos")
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[1], nil
}

func main() {
	nums := []int{5, 10, 3, 8, 15}
	secondLargest, err := findSecondLargest(nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("O segundo maior valor é:", secondLargest)
	}
}
