package main

import "fmt"

func ysf(arr []int, num int) []int {
    
    var ret []int

    i := 0
    j := 0
    for len(arr) > 0 {
        i++ 
        if j == len(arr) {
            j = 0
        }
        if i == num {
            i = 0
            ret = append(ret, arr[j])
            arr = append(arr[0:j], arr[j+1:]...)
            j--
        }
        j++
    }

    return ret
}

func main() {
    fmt.Println(ysf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6))
    // 6 3 1 9 2 5 4 8 7

    // 6 3 1 9 2 5 4 8 7 
}
