// leetcode 213
package main

import "fmt"

func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    if n == 2 && nums[0] > nums[1] {
        return nums[0]
    } else if n == 2{
        return nums[1]
    }

    with := nums[0] + _rob(nums[2:n-1])
    without := _rob(nums[1:])

    if with > without {
        return with
    }
    return without
}

func _rob(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    for i:=1; i<len(nums); i++ {
        if i == 1 {
            dp[i] = max(dp[0], nums[1])
            continue
        }
        dp[i] = max(dp[i-2]+nums[i], dp[i-1])
    }

    ret := nums[0]
    for i:=1; i<len(dp); i++ {
        if ret < dp[i] {
            ret = dp[i]
        }
    }
    return ret
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(rob([]int{1, 2, 3, 1}))
    fmt.Println(rob([]int{1, 2, 1, 1}))
    fmt.Println("---")
    fmt.Println(_rob([]int{2, 1, 1}))
}

