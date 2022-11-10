// A permutation of an array of integers is an arrangement of its members into a sequence or linear order.
//
// For example, for arr = [1,2,3], the following are all the permutations of arr: [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
// The next permutation of an array of integers is the next lexicographically greater permutation of its integer. 
// More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, 
// then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).
//
// For example, the next permutation of arr = [1,2,3] is [1,3,2].
// Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
// While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
// Given an array of integers nums, find the next permutation of nums.
//
// The replacement must be in place and use only constant extra memory.
//
//
//
// Example 1:
//
// Input: nums = [1,2,3]
// Output: [1,3,2]
// Example 2:
// Input: nums = [3,2,1]
// Output: [1,2,3]
// Example 3:
// Input: nums = [1,1,5]
// Output: [1,5,1]
// Example 4:
// Input: nums = [1,3,2] [1, 5, 4, 3]
// Output: [2,1,3]
//
// Constraints:
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100


package main

import "fmt"

func nextPermutation(nums []int) []int {
    k := len(nums) - 1
    for k > 0 {
        if nums[k] > nums[k-1] {
            break
        }
        k--
    }
    if k == 0 {
        // 321
        for i:=0; i<len(nums)/2; i++ {
            nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
        }
        return nums
    }
    m := len(nums) - 1
    for m > k {
        if nums[m] > nums[k-1] {
            break
        }
        m--
    }
    nums[k-1], nums[m] = nums[m], nums[k-1]
    for i:=k; i<len(nums)-1; i++ {
        for j:=i+1; j<len(nums); j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }

    return nums
}

func main() {
    fmt.Println("[1,2,3] : ", nextPermutation([]int{1, 2, 3}))
    fmt.Println("[1,3,2] : ", nextPermutation([]int{1, 3, 2}))
    fmt.Println("[3,2,1] : ", nextPermutation([]int{3, 2, 1}))
    fmt.Println("[4,3,2,1] : ", nextPermutation([]int{4, 3, 2, 1}))
    fmt.Println("[10,9,8,7,6,5,4,3,2,1] : ", nextPermutation([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
    fmt.Println("[1,1,5] : ", nextPermutation([]int{1, 1, 5}))
    fmt.Println("[4,3,5,1] : ", nextPermutation([]int{4, 3, 5, 1}))
    fmt.Println("[4,3,5,1] : ", nextPermutation([]int{100,99,98,97,96,95,94,93,92,91,90,89,88,87,86,85,84,83,82,81,80,79,78,77,76,75,74,73,72,71,70,69,68,67,66,65,64,63,62,61,60,59,58,57,56,55,54,53,52,51,50,49,48,47,46,45,44,43,42,41,40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1}))
}
