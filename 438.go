package main

import "fmt"

func findAnagrams(s string, p string) []int {
    n := len(s)
    if n < len(p) || n == 0 {
        return []int{0}
    }
    
    ret := []int{}
    
    pm := make(map[byte]int)
    for i:=0; i<len(p); i++ {
        pm[p[i]]++
    }
    
    sm := make(map[byte]int)
    l, r := 0, 0
    for r < n {
        
        // add r to window
        sm[s[r]]++
        
        // status check
        // s[r] not in pm
        if pm[s[r]] == 0 {
            fmt.Println("s[r]: ", string(s[r]), "not exists in p, r: ", r, " l:", l)
            r++
            l = r
            sm = make(map[byte]int)
            continue
        }
        
        // status check
        // s[r] in pm & sm[s[r]] <= pm[s[r]]  
        if sm[s[r]] == pm[s[r]] && em(sm, pm) {
            fmt.Println(" === s[r]: ", string(s[r]), ", r: ", r, " l:", l)
            ret = append(ret, l)
            sm[s[l]]--
            l++
        }
        
        // status check
        // sm[s[r]] > pm[s[r]]
        for sm[s[r]] > pm[s[r]] {
            fmt.Println(" >>> sm:", sm, " pm: ", pm, " s[r]: ", s[r], ", r: ", r, " l:", l)
            sm[s[l]]--
            l++
        }
        r++
    }
    return ret
}

func em(sm map[byte]int, pm map[byte]int) bool {
    if len(sm) != len(pm) {
        return false
    }
    
    for k, v := range(sm){
        if v != pm[k] {
            return false
        }
    }
    return true
}

func main() {
    // fmt.Println(findAnagrams("cbaebabacd", "abc"))
    fmt.Println(findAnagrams("abaacbabc", "abc"))
}
