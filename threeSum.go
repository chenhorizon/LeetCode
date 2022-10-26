package main

import "fmt"
import "strings"

func bubbleSort(nums []int) []int {
    for i:=0; i<len(nums)-1; i++ {
        for j:=i+1; j<len(nums); j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
    return nums
}

func quickSort(nums []int) []int{
    if len(nums) <=1 {
        return nums
    }
    partition := getPartition(nums)
    quickSort(nums[0:partition])
    quickSort(nums[partition+1:])
    
    return nums
}

func getPartition(nums []int) int {
    pv := nums[0]
    index := 1
    for i:=1; i<len(nums); i++ {
        if pv > nums[i] {
            nums[i], nums[index] = nums[index], nums[i]
            index++
        }
    }
    nums[0], nums[index-1] = nums[index-1], nums[0]

    return index-1
}

func uniqSortedArray(arr [][]int) [][]int {
    var ret [][]int
    if len(arr) < 1 {
        return ret
    }
    for i:=0; i<len(arr)-1; i++ {
        ii1 := false
        if len(arr[i]) == len(arr[i+1]) {
            ii1 = true
            for k:=0; k<len(arr[i]); k++ {
                if arr[i][k] != arr[i+1][k] {
                    ii1 = false
                    break
                }
            }
        }
        if ii1 == false {
            ret = append(ret, arr[i])
        }
    }
    ret = append(ret, arr[len(arr)-1])

    return ret
}

func uniqSlice(arr [][]int) [][]int {
    retMap := make(map[string][]int)
    var ret [][]int
    if len(arr) <= 1 {
        return arr
    }
    for i:=0; i<len(arr); i++ {
        strArr := strings.Trim(strings.Join(strings.Split(fmt.Sprint(arr[i]), " "), ""), "[]")
        retMap[strArr] = arr[i]
    }

    for _, v := range(retMap) {
        ret = append(ret, v)
    }
        
    return ret
}

func threeSum(nums []int) [][]int {

    nums = elementTimes(nums, 3)
    nums = bubbleSort(nums)
    nums = quickSort(nums)

    var ret [][]int
    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            for k:=j+1; k<len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 0 {
                    ret = append(ret, []int{nums[i], nums[j], nums[k]})
                }
            }
        }
    }
    ret = uniqSlice(ret)

    return ret
}

func threeSumNew(nums []int) [][]int {
    var ret [][]int

    var pNums []int 
    var nNums []int
    var zNums []int
    for i:=0; i<len(nums); i++ {
        if nums[i] < 0 {
            nNums = append(nNums, nums[i])
        } else if nums[i] == 0 {
            zNums = append(zNums, nums[i])
        } else {
            pNums = append(pNums, nums[i])
        }
    }

    if len(zNums) >= 3 {
        ret = append(ret, []int{0, 0, 0})
    }

    nNums = quickSort(nNums)
    pNums = quickSort(pNums)

    for i:=0; i<len(nNums); i++ {
        for j:=i+1; j<len(nNums); j++ {
            for k:=0; k<len(pNums); k++ {
                r := nNums[i] + nNums[j] + pNums[k]
                if r == 0 {
                    ret = append(ret, []int{nNums[i], nNums[j], pNums[k]})
                } else if r > 0 {
                    break
                }
            }
        }
    }

    for i:=0; i<len(pNums); i++ {
        for j:=i+1; j<len(pNums); j++ {
            for k:=0; k<len(nNums); k++ {
                r := pNums[i] + pNums[j] + nNums[k]
                if r == 0 {
                    ret = append(ret, []int{pNums[i], pNums[j], nNums[k]})
                } else if r > 0 {
                    break
                }
            }
        }
    }

    if len(zNums) > 0 {
        for i:=0; i<len(nNums); i++ {
            for j:=0; j<len(pNums); j++ {
                r := nNums[i] + pNums[j]
                if r == 0 {
                    ret = append(ret, []int{nNums[i], 0, pNums[j]})
                } else if r > 0 {
                    break
                }
            }
        }
    }

    ret = uniqSlice(ret)
    return ret
}

func elementTimes(arr []int, times int) []int {
    var ret []int
    retMap := make(map[int]int)
    for _, v := range(arr) {
        if retMap[v] < times {
            retMap[v] = retMap[v] + 1
        }
    }

    for k, v := range(retMap) {
        for i:=0; i<v; i++ {
            ret = append(ret, k)
        }
    }

    return ret
}

func main() {
    fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
    fmt.Println(threeSumNew([]int{-1, 0, 1, 2, -1, -4}))
    fmt.Println(threeSum([]int{0, 0, 0}))
    fmt.Println(threeSumNew([]int{0, 0, 0}))
}
