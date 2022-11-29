// leetcode75
package main

import "fmt"

func sortColors(nums []int)  {

    n := len(nums)
    if n == 1 {
        return
    }

    r := n-1
    for i:=0; i<n; i++ {
        r = n - 1
        if nums[i] == 0 {
            continue
        } else if nums[i] == 1 {
            for r > i && nums[r] != 0 {
                r--
            }
            if r ==i {
                continue
            }
            nums[i], nums[r] = nums[r], nums[i]
        } else if nums[i] == 2 {
            for r > i && nums[r] == 2 {
                r--
            }
            if r == i {
                break
            }
            nums[i], nums[r] = nums[r], nums[i]
            if nums[i] == 1 {
                i--
            }
        }
    }
}

func main() {
    nums := []int{2,0,2,1,1,0}
    //nums = []int{1,2,1}
    sortColors(nums)
    fmt.Println("nums: ", nums) 
}
