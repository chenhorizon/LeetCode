// leetcode128
package main

import "fmt"
import "sort"

func longestConsecutive(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return n
    }

    sort.Ints(nums)
    max := 1
    min := 1
    for i:=1; i<n; i++ {
        if nums[i] == nums[i-1] + 1 {
            min++
            continue
        } else if nums[i] == nums[i-1] {
            continue
        } else {
            if min > max {
                max = min
            }
            min = 1
        }
    }
    if min > max {
        max = min
    }

    return max
}

func main() {
    fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
}
