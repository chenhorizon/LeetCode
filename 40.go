// leetcode 40
package main

import "fmt"
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    n := len(candidates)
    if n == 1 {
        if target == candidates[0] {
            return [][]int{{target}}
        }
        return nil
    }

    var ret [][]int
    sort.Ints(candidates)

    if target < candidates[0] {
        return nil
    }

    for i:=0; i<n-1; i++ {
        if i>0 {
            continue
        }

        if candidates[i] > target {
            continue
        }

        if candidates[i] == target {
            return [][]int{{target}}
        }

        with := combinationSum2(candidates[i+1:], target-candidates[i])
        j:=i+1
        for j <= n-1 && candidates[j] == candidates[i] {
            j++
        }
        var without [][]int
        if j <= n-1 {
            without = combinationSum2(candidates[j:], target)
        }

        if with != nil {
            for _, v := range(with) {
                v = append(v, candidates[i])
                ret = append(ret, v)
            }
        }
        if without != nil {
            for _, v := range(without) {
                ret = append(ret, v)
            }
        }
    }

    return ret
}

func main() {
    fmt.Println(combinationSum2([]int{2,3,5,6}, 8))
    fmt.Println(combinationSum2([]int{8}, 8))
    fmt.Println(combinationSum2([]int{8}, 9))
    fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
    fmt.Println(combinationSum2([]int{1,1,2,5,6,7,10}, 8))
}
