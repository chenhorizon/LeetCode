package main

import "fmt"

func bubbleSort(arr []int) []int{
    length := len(arr)
    for i:=0; i<length-1; i++ {
        for j:=i+1; j<length; j++ {
            if(arr[i] > arr[j]) {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }

    return arr
}

func selectionSort(arr []int) []int{
    length := len(arr)
    for i:=0; i<length-1; i++ {
        index := i
        for j:=i+1; j<length; j++ {
            if(arr[j] < arr[index]) {
                index = j
            }
        }
        arr[i], arr[index] = arr[index], arr[i]
    }

    return arr
}

func insertionSort(arr []int) []int {
    length := len(arr)

    i, j := 0, 0
    for i=1; i<length; i++ {
        c := arr[i]
        for j=i-1; j>=0 && arr[i]<arr[j]; j-- {
            arr[j+1] = arr[j]
        }
        arr[j+1] = c
    }

    return arr
}

func mergeSort(arr []int) []int{
    length := len(arr)
    if length < 2 {
        return arr
    }
    m := length / 2
    leftArr := arr[0:m]
    rightArr := arr[m:]
    ret := _merge(mergeSort(leftArr), mergeSort(rightArr))
    return ret
}

func _merge(lArr []int, rArr []int) []int{
    var ret []int
    for len(lArr) != 0 && len(rArr) != 0 {
        if lArr[0] <= rArr[0] {
            ret = append(ret, lArr[0])
            lArr = lArr[1:]
        } else {
            ret = append(ret, rArr[0])
                rArr = rArr[1:]
        }
    }
    for len(lArr) > 0 {
        ret = append(ret, lArr[0])
        lArr = lArr[1:]
    }
    for len(rArr) > 0 {
        ret = append(ret, rArr[0])
        rArr = rArr[1:]
    }

    return ret
}

func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    pv := arr[0]
    pi := 1
    for i:=1; i<len(arr); i++ {
        if pv > arr[i] {
            arr[i], arr[pi] = arr[pi], arr[i]
            pi++
        }
    }
    arr[pi-1], arr[0] = arr[0], arr[pi-1]
    quickSort(arr[0:pi-1])
    quickSort(arr[pi:])

    return arr
}

func main() {
    fmt.Println(bubbleSort([]int{3, 2, 4}))
    fmt.Println(selectionSort([]int{3, 2, 4}))
    fmt.Println(insertionSort([]int{3, 2, 4}))
    fmt.Println(mergeSort([]int{3, 2, 4, 5, 1, 9, 0}))
    fmt.Println(quickSort([]int{3, 2, 4, 5, 1, 9, 0}))
    fmt.Println(quickSort([]int{3, 2, 4, 5, 1, 3, 0}))
    arr := []int{3, 2, 4, 5, 1, 3, 0}
    fmt.Println("---")
    fmt.Println(quickSort(arr))
    fmt.Println(arr)
}
