package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
    minSize := len(nums)

    l, r := 0, 0
    sum := 0
    for r < len(nums) {
        sum += nums[r]
        for sum >= target && l <= r {
            if minSize > r - l + 1 {
                minSize = r - l + 1
            }
            sum = sum - nums[l]
            l++
        }
        r++
    }

    if l == 0 && r == len(nums) {
        minSize = 0
    }

    return minSize
}

func main() {
    // fmt.Println(minSubArrayLen(7, []int{2,3,1,2,4,3}))
    fmt.Println(minSubArrayLen(7, []int{4,3}))
}
