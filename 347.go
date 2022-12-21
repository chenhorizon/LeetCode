// leetcode 347

import "sort"

type Node struct {
    Value int
    Times int
}

func topKFrequent(nums []int, k int) []int {

    nm := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        nm[nums[i]]++
    }

    var ns []Node
    for k, v := range nm {
        n := Node{k, v}
        ns = append(ns, n)
    }

    sort.Slice(ns, func(i, j int) bool {
        return ns[i].Times > ns[j].Times
    })

    var unums []int
    for i:=0; i<k; i++ {

        unums = append(unums, ns[i].Value)
    }

    return unums
}
