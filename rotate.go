package main

import "fmt"

func rotate(matrix [][]int) {
    n := len(matrix)

    fmt.Println("Before: ", matrix)
    r1, r2, r3, r4 := 0, 0, 0, 0
    c1, c2, c3, c4 := 0, 0, 0, 0
    row := 0
    for n > 0 {
        for i:=0; i<n-1; i++ {
            r1, c1 = row, row+i
            r2, c2 = c1 + n -1, n-1+row
            if r2 >= row+n-1 {
                r2 = row + (r2 % (row+n-1))
            }
            r3, c3 = n-1+row, r2 + n - 1
            if c3 >= row+n-1 {
                c3 = row+n-1 - (c3 % (row+n-1))
            }
            r4, c4 = c3 - n + 1, row
            if r4 <= row {
                r4 = row+n-1 + (r4 - row)
            }
            fmt.Println("row:", row, "col:", row+i)
            fmt.Println("1:", r1, c1)
            fmt.Println("2:", r2, c2)
            fmt.Println("3:", r3, c3)
            fmt.Println("4:", r4, c4)
            
            matrix[r1][c1], matrix[r2][c2], matrix[r3][c3], matrix[r4][c4] = matrix[r4][c4], matrix[r1][c1], matrix[r2][c2], matrix[r3][c3]
        }
        n = n - 2
        row++
    }

    fmt.Println("After: ", matrix)
}

func main() {

    rotate([][]int{{1, 2}, {3, 4}})
    rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
    rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})

}
