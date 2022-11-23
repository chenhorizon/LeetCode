//leetcode27
package main

func removeElement(nums []int, val int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    for i:=0; i<n; i++ {
        if nums[i] == val {
            if i == n-1 {
                return n-1
            }
            for j:=i+1; j<n; j++ {
                nums[j-1] = nums[j]
            }
            i--
            n--
        }
    }
    return n
}
