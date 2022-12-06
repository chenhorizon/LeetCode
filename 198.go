// leetcode198

package main

import "fmt"

func rob(nums []int) int {
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
