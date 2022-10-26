// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// Example 1:
//
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// Example 2:
//
// Input: n = 1
// Output: ["()"]
//  
//
// Constraints:
//
// 1 <= n <= 8


package main

import "fmt"

func generateParenthesis(n int) []string {

    retArr := []string{"("}
    for len(retArr) > 0  {
        k :=0
        v := retArr[k]
        lNum := calculateElementNum(v, '(')
        rNum := calculateElementNum(v, ')')
        if rNum == n {
            break
        }
        nextP := ""
        if lNum > rNum {
            if n-lNum > 0 {
                nextP = "("
                retArr = append(retArr, v + nextP)
            }
            if n-rNum > 0 {
                nextP = ")"
                retArr = append(retArr, v + nextP)
            }
        }else if lNum == rNum {
            if n-lNum > 0 {
                nextP = "("
                retArr = append(retArr, v + nextP)
            }
        }else {
            break;
        }

        retArr = retArr[1:]
    }

    return retArr
}

func calculateElementNum(target string, e rune) int {
    ret := 0
    for _, v := range(target) {
        if v == e {
            ret++
        }
    }
    return ret
}

func main() {
    fmt.Println(generateParenthesis(1))
    fmt.Println(generateParenthesis(2))
    fmt.Println(generateParenthesis(3))
    fmt.Println(generateParenthesis(4))
    fmt.Println(generateParenthesis(5))
    fmt.Println(generateParenthesis(6))
    fmt.Println(generateParenthesis(7))
    fmt.Println(generateParenthesis(8))
}
