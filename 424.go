package main

import "fmt"

func characterReplacement(s string, k int) int {

    n := len(s)
    if n == 0 { return 0 }
    if n <=k+1 { return n }

    ret := 0
    l := 0
    sm := make(map[byte]int)
    for r, _ := range(s) {
        sm[s[r]]++

        if (r-l+1) - maxFrequency(sm) > k {
            sm[s[l]]--
            l++
        }

        if ret < r-l+1 {
            ret = r-l+1
        }
    }

    return ret
}

func maxFrequency(sm map[byte]int) int {
    max := 0
    for _, v := range(sm) {
        if max < v {
            max = v
        }
    }

    return max
}

func main() {
    fmt.Println(characterReplacement("AABABBA", 1))
}
