package main 

import "fmt"

func uniquePaths(m int, n int) int {
    if m == 1 || n == 1 {
        return 1
    }

    r, rn, c, cn := 0, 0, 0, 0
    if r + 1 < m {
        rn = uniquePaths(m, n-1)
    }
    if c + 1 < n {
        cn = uniquePaths(m-1, n)
    }

    return rn + cn

}

func main() {
    fmt.Println("m:3 n:2 :::", uniquePaths(3, 2), " unique paths.")
    fmt.Println("m:3 n:7 :::", uniquePaths(3, 7), " unique paths.")
}
