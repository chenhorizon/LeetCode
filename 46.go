// leetcode46
package main

import "fmt"

func permute(nums []int) [][]int {
    n := len(nums)
    if n == 1 {
        return [][]int{{nums[0]}}
    }

    var ret [][]int
    for i:=0; i<n; i++ {
        var remains []int
        p := nums[i]
        if i == 0 {
            remains = append(remains, nums[1:]...)
        } else if i == n-1 {
            remains = append(remains, nums[0:n-1]...)
        } else {
            remains = append(remains, nums[0:i]...)
            remains = append(remains, nums[i+1:]...)
        }
        pr := permute(remains)
        for _, vr := range(pr) {
            vr = append(vr, p)
            ret = append(ret, vr)
        }
    }

    return ret
}

func main() {
    nums := []int{1, 2, 3}
    fmt.Println("Premute ", nums, ": ", permute(nums))
}
