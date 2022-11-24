// leetcode47
package main

import "fmt"

func permuteUnique(nums []int) [][]int {
    
    if len(nums) == 1 {
        return [][]int{{nums[0]}}
    }

    var ret [][]int
    n := len(nums)
    for i:=0; i<n; i++ {
        p := nums[i]
        var remains []int
        if i == 0 {
            remains = append(remains, nums[1:]...)
        } else if i == n-1 {
            remains = append(remains, nums[:i]...)
        } else {
            remains = append(remains, nums[0:i]...)
            remains = append(remains, nums[i+1:]...)
        }

        iret := permuteUnique(remains)
        for _, item := range(iret) {
            item = append(item, p)
            ret = append(ret, item)
        }
    }
    ret = unique(ret)

    return ret
}

func unique(nums [][]int) [][]int {
    hm := make(map[string][]int)
    for _, item := range(nums) {
        k := ""
        for _, v := range(item) {
            if v < 0 {
                v = -1 * v
                k = k + ":" + "-" + string(v)
            } else {
                k = k + ":" + string(v)
            }
        }
        hm[k] = item
    }

    var ret [][]int
    for _, item := range(hm) {
        ret = append(ret, item)
    }

    return ret
}

func main() {
    nums := []int{1, 1, 1, 3}
    fmt.Println("permuteUnique of nums ", nums, ": ", permuteUnique(nums))
    nums = []int{-1,-2,0,-1,1,0,1}
    fmt.Println("permuteUnique of nums ", nums, ": ", permuteUnique(nums))
}
