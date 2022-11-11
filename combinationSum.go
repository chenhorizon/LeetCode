// Given an array of distinct integers candidates and a target integer target, 
// return a list of all unique combinations of candidates where the chosen numbers sum to target. 
// You may return the combinations in any order.
//
// The same number may be chosen from candidates an unlimited number of times. 
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.
//
// The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations 
// for the given input.
//  
//
// Example 1:
//
// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.
// Example 2:
//
// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]
// Example 3:
//
// Input: candidates = [2], target = 1
// Output: []
//  
//
// Constraints:
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40

package main

import "fmt"
import "sort"

func combinationSum(candidates []int, target int) [][]int {
    ret := [][]int{}
    fRet := [][]int{}

    sort.Ints(candidates)
    times := target / candidates[0] + 1

    for times > 0 {
        lenRet := len(ret)
        for i:=0; i<len(candidates); i++ {
            if lenRet == 0 {
                ret = append(ret, []int{candidates[i]})
                if candidates[i] == target {
                    fRet = append(fRet, []int{candidates[i]})
                    fmt.Println("zzz:", fRet)
                }
                continue
            }
            for j:=0; j<lenRet; j++ {
                if ret[j][len(ret[j])-1] > candidates[i] {
                    continue
                }
                ret = append(ret, append(ret[j], candidates[i]))
                sum := 0
                for k:=0; k<len(ret[len(ret)-1]); k++ {
                    sum = sum + ret[len(ret)-1][k]
                }
                if sum == target {
                    fmt.Println(ret[len(ret)-1])
                    cpRetItem := make([]int,len(ret[len(ret)-1]))
                    copy(cpRetItem, ret[len(ret)-1][:])
                    fRet = append(fRet, cpRetItem)
                }
            }
        }

        ret = ret[lenRet:]
        times--
    }

    fmt.Println("eee:", fRet)
    return fRet
}

func main() {
    fmt.Println("candidates = [2,3,5], target = 8:", combinationSum([]int{2, 3, 5}, 8))
    fmt.Println("candidates = [2,3,6], target = 9:", combinationSum([]int{2, 3, 6}, 9))
    fmt.Println("candidates = [2,3,6,7], target = 7:", combinationSum([]int{2, 3, 6, 7}, 7))
    fmt.Println("candidates = [7,3,2], target = 18:", combinationSum([]int{7, 3, 2}, 18))
    fmt.Println("candidates = [2,3,7], target = 18:", combinationSum([]int{2, 3, 7}, 18))
}
