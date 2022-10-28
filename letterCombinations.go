// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. 
// Return the answer in any order.
//
// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
// Example 1:
//
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Example 2:
//
// Input: digits = ""
// Output: []
// Example 3:
//
// Input: digits = "2"
// Output: ["a","b","c"]
//  
//
// Constraints:
//
// 0 <= digits.length <= 4
// digits[i] is in the range ['2', '9'].

package main

import "fmt"
import "strings"

func letterCombinations(digits string) []string {
    var ret []string
    dict := make(map[string][]string)

    dict["2"] = []string{"a", "b", "c"}
    dict["3"] = []string{"d", "e", "f"}
    dict["4"] = []string{"g", "h", "i"}
    dict["5"] = []string{"j", "k", "l"}
    dict["6"] = []string{"m", "n", "o"}
    dict["7"] = []string{"p", "q", "r", "s"}
    dict["8"] = []string{"t", "u", "v"}
    dict["9"] = []string{"w", "x", "y", "z"}

    for _, vDigits := range(strings.Split(digits, "")) {
        if len(ret) == 0 {
            for _, vDict := range(dict[vDigits]) {
                ret = append(ret, vDict)
            }
            continue
        }
        lRet := len(ret)
        for _, vRet := range(ret) {
            for _, vDict := range(dict[vDigits]) {
                ret = append(ret, vRet+vDict)
            }
        }
        ret = ret[lRet:]
    }

    return ret
}

func main() {
    fmt.Println(letterCombinations("234"))
}
