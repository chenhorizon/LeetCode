// leetcode26
package main

func removeDuplicates(nums []int) int {
    
    if len(nums) == 1 {
        return 1
    }
    i := 0
    p := i + 1
    n := len(nums)
    for i < n {
        for p < n && nums[p] == nums[p-1] {
            p++
        }
        if p == n {
            return i+1
        }
        for j:=1; j<=n-p; j++ {
            nums[i+j] = nums[p+j-1]
        }
        n = n - p + i + 1
        i++
        p = i + 1
    }
    return n
}
