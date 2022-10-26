package main

import "strings"
import "fmt"

func longestPalindromic(s string) string {

    if(len(s) <= 1) {
        return s
    }
    // 求索引
    sArr := strings.Split(s, "")
    sMap := make(map[string][]int)
    for sIndex, sItem := range sArr {
        sMap[sItem] = append(sMap[sItem], sIndex)
    }

    // 求Palindromic
    max := 0
    ret := ""
    for sItem, sIndexes := range sMap {
        if(len(sIndexes) == 1) {
            if(max < 1) {
                max = 1
                ret = sItem
            }
            continue
        }
        for i:=0; i<len(sIndexes); i++ {
            for j:=i+1; j<len(sIndexes); j++ {
                pss := sArr[sIndexes[i]:sIndexes[j]+1]
                if( (max < len(pss)) && isPalindromicArray(pss)) {
                    max = len(pss)
                    ret = strings.Join(pss, "")
                }
            }
        }
    }

    return ret
}

func isPalindromicArray(sArr []string) bool {
    start := 0
    end := len(sArr) - 1

    if(end == 0 || end == -1) {
        return true
    }

    for start < end {
        if(sArr[start] != sArr[end]) {
            return false
        }
        start++
        end--
    }
    return true
}

func isPalindromic(s string) bool {
    start := 0
    end := len(s) - 1
    sArr := strings.Split(s, "")

    if(end == 0 || end == -1) {
        return true
    }

    for start < end {
        if(sArr[start] != sArr[end]) {
            return false
        }
        start++
        end--
    }
    return true
}

func main() {
    fmt.Println("abccba: " + longestPalindromic("abccba"))
    fmt.Println("abab: " + longestPalindromic("abab"))
    fmt.Println("ababa: " + longestPalindromic("ababa"))
    fmt.Println("aaa: " + longestPalindromic("aaa"))
    fmt.Println("aa: " + longestPalindromic("aa"))
    fmt.Println("a: " + longestPalindromic("a"))
}
