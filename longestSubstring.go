package main

import "strings"
import "fmt"
import "strconv"

func longestSubstring(s string) int {

    max := 0
    currentMax := 0
    sArr := strings.Split(s, "")
    ret := ""
    for i:=0; i<len(sArr); i++ {
        ret = sArr[i]
        for j:=i+1; j<len(sArr); j++ {
            idx := strings.LastIndex(ret, sArr[j])
            if(idx == -1) {
                ret = ret + sArr[j]
                currentMax = len(ret)
                if(currentMax > max) {
                    max = currentMax
                }
            }else {
                i = i + idx
                break
            }
        }
        currentMax = len(ret)
        if(currentMax > max) {
            max = currentMax
        }
    }
    return max
}

func ls(s string) int {
    dict := make(map[string][]int)
    maxLen := 0
    nowLen := 0
    for k, v := range(s) {
        vs := strconv.QuoteRune(v)
        if len(dict[vs]) <= 0 {
            nowLen++
            if nowLen > maxLen {
                maxLen = nowLen
            }
            dict[vs] = append(dict[vs], k)
        } else {
            nowLen = 1
            for kd := range(dict) {
                delete(dict, kd)
            }
            dict[vs] = append(dict[vs], k)
        }
    }
    
    return maxLen
}

func main() {
    fmt.Println(longestSubstring("bacabcabc"))
    fmt.Println(longestSubstring("kwwekd"))
    fmt.Println(longestSubstring("abcde"))
    fmt.Println(longestSubstring(" "))
    fmt.Println(longestSubstring("a"))
    fmt.Println(longestSubstring("aab"))
    fmt.Println(longestSubstring("aa"))
    fmt.Println("---")
    fmt.Println(ls("bacabcabc"))
    fmt.Println(ls("kwwekd"))
    fmt.Println(ls("abcde"))
    fmt.Println(ls(" "))
    fmt.Println(ls("a"))
    fmt.Println(ls("aab"))
    fmt.Println(ls("aa"))
}
