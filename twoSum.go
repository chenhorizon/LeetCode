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

func vkArray(arr []int) map[int][]int{
    ret := make(map[int][]int)
    for k, v := range(arr) {
        ret[v] = append(ret[v], k)
    }

    return ret
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

func twoSum(nums []int, target int) []int {
    vkNums := vkArray(nums)
    nums = mergeSort(nums)
    length := len(nums)
    for i:=0; i<length-1; i++{
        for j:=i+1; j<length; j++ {
            if nums[j] + nums[i] == target {
                // fmt.Println(nums[i], i)
                // fmt.Println(nums[j], j)
                ik := vkNums[nums[i]][0]
                if len(vkNums[nums[i]]) > 1 {
                    vkNums[nums[i]] = vkNums[nums[i]][1:]
                }
                jk := vkNums[nums[j]][0]
                return []int{ik, jk}
            } else if nums[j] + nums[i] > target {
                break
            }
        }
    }

    return []int{0, 0}
}


func main() {
    fmt.Println(twoSum([]int{2, 2, 4}, 7))
    fmt.Println(twoSum([]int{3, 2, 4}, 7))
    fmt.Println(twoSum([]int{2, 2, 4}, 6))
    fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
