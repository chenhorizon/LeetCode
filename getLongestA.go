package main

import "fmt"

func ga(arr []string) int {
    
    if len(arr) == 0 {
        return 0
    }

    maxLen := 0

    var bIndex []int
    for k, v := range(arr) {
        if v == "B" {
            bIndex = append(bIndex, k)
        }
    }

    if len(bIndex) == 0 {
        return len(arr) - 1
    }

    for k, v := range(bIndex) {
        beforeB := 0
        afterB := 0
        if k == 0 {
            beforeB = -1
        } else {
            beforeB = bIndex[k-1]
        }
        if k + 1 == len(bIndex) {
            afterB = len(arr)
        }else {
            afterB = bIndex[k+1]
        }
        currentLen := v - beforeB -1 + afterB - v -1
        if currentLen > maxLen {
            maxLen = currentLen
        }
    }

    return maxLen
}

func nga(arr []string) int {
    oneB := false
    hasB := false
    index := 0
    maxLen := 0
    nowLen := 0

    for index < len(arr) {
        if arr[index] == "A" {
            nowLen++
            if nowLen > maxLen {
                maxLen = nowLen
            }

        } else if arr[index] == "B" {
            hasB = true
            if oneB == false && nowLen != 0 {
                oneB = true
            } else if oneB == true{
                if nowLen > maxLen {
                    maxLen = nowLen
                }
                nowLen = 0
                oneB = false
            }
        }
        index++
    }
    if (hasB == false) {
        maxLen--
    }

    return maxLen
}

func main() {
    fmt.Println("222330");
    fmt.Println(ga([]string{"B", "B", "A", "A"}))
    fmt.Println(ga([]string{"B", "A", "B", "A"}))
    fmt.Println(ga([]string{"A", "B", "A", "B"}))
    fmt.Println(ga([]string{"A", "A", "B", "A"}))
    fmt.Println(ga([]string{"A", "A", "A", "A"}))
    fmt.Println(ga([]string{"B", "B", "B", "B"}))
    fmt.Println("---")
    fmt.Println(nga([]string{"B", "B", "A", "A"}))
    fmt.Println(nga([]string{"B", "A", "B", "A"}))
    fmt.Println(nga([]string{"A", "B", "A", "B"}))
    fmt.Println(nga([]string{"A", "A", "B", "A"}))
    fmt.Println(nga([]string{"A", "A", "A", "A"}))
    fmt.Println(nga([]string{"B", "B", "B", "B"}))

    fmt.Println(fmt.Sprint([]int{1, 2, 3}))
}
