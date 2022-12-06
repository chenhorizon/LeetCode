// leetcode39
package main

import "fmt"
import "sort"

func combinationSum(candidates []int, target int) [][]int {

    n := len(candidates)
    if n == 1 && target % candidates[0] == 0 {
        var r []int
        for i:=0; i<target/candidates[0]; i++ {
            r = append(r, candidates[0])
        }
        return [][]int{r}
    } else if n == 1 || n == 0 {
        return [][]int{}
    }
    sort.Ints(candidates)
    if target < candidates[0] {
        return [][]int{}
    } else if target == candidates[0] {
        return [][]int{{candidates[0]}}
    }

    var ret [][]int
    p := candidates[i]
    withP := combinationSum(candidates[i:], target-p)
    for _, item := range(withP) {
        item = append(item, p)
        ret = append(ret, item)
    }

    withoutP := combinationSum(candidates[i+1:], target)
    for _, item := range(withoutP) {
        ret = append(ret, item)
    }

    return ret
}

func main() {
    candidates := []int{2, 3, 6, 7}
    target := 7
    fmt.Println("candidates: ", candidates, " target: ", target, " : ")
    fmt.Println(combinationSum(candidates, target))
}
