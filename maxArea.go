package main

import "fmt"

func maxArea(height []int) int {

    max := 0
    lHeight := len(height)
    for i:=0; i<lHeight-1; i++ {
        for i>0 && height[i] < height[i-1] {
            i++
        }
        for j:=i+1; j<lHeight; j++ {
            for j+1<lHeight && height[j] < height[j+1] {
                j++
            }
            area := 0
            if height[i] < height[j] {
                area = (j-i) * height[i]
            }else {
                area = (j-i) * height[j]
            }
            if area > max {
                max = area
            }
        }
    }

    return max
}

func maxAreaNew(height []int) int {
    max := 0
    area := 0
    l := 0 
    r := len(height) - 1

    for l < r {

        if height[l] < height[r] {
            area = (r-l) * height[l]
        }else {
            area = (r-l) * height[r]
        }
        if area > max {
            max = area
        }

        if height[l] < height[r] {
            l++
        }else {
            r--
        }
    }

    return max
}

func main() {
    fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
    fmt.Println(maxArea([]int{1, 1}))

    fmt.Println(maxAreaNew([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
    fmt.Println(maxAreaNew([]int{1, 1}))
}
