package main

import (
	"fmt"
)

// slices: It is a dynamic array, We use slices when we don't know size of data
// + useful methods are available in slices
func main() {
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)

	// capacity -> maxium numbers of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)
	// nums = append(nums, 10)
	// nums = append(nums, 7)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// var nums = make([]int, 2, 5)

	// nums[0] = 3
	// nums[1] = 4
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// another aproach to make slices

	// nums := []int{}

	// nums = append(nums, 1)
	// nums = append(nums, 2)

	// copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)

	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)

	// fmt.Println(len(nums2))
	// fmt.Println(nums, nums2)

	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[0:])
	// fmt.Println(nums[1:2])

	// slice

	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(nums)
}
