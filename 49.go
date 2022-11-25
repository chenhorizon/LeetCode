// leetcode49
package main

func groupAnagrams(strs []string) [][]string {
    ls := len(strs)
    ret := make(map[string][]string)
    if ls == 0 {
        return [][]string{{""}}
    }
    if ls == 1 {
        return [][]string{{strs[0]}}
    }

    for _, v := range(strs) {
        sg := sig(v)
        if item, has := ret[sg]; has {
            ret[sg] = append(ret[sg], v)
            fmt.Println(item)
        } else {
            ret[sg] = []string{v}
        }
    }

    var r [][]string
    for _, v := range(ret) {
        r = append(r, v)
    }

    return r
}

func sig(s string) string {
    ls := len(s)
    if ls == 0 || ls == 1 {
        return s
    }
    bs := []byte(s)
    for i:=0; i<ls; i++ {
        for j:=i+1; j<ls; j++ {
            if bs[i] < bs[j] {
                bs[i], bs[j] = bs[j], bs[i]
            }
        }
    }
    s = string(bs[:])
    return s
}
