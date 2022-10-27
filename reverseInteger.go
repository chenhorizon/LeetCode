package main

import "fmt"
import "math"
import "strings"
import "strconv"


func reversePro(x int) int {
	signMultiplier := 1
	if x < 0 {
		signMultiplier = -1
		x = signMultiplier * x
	}

	var result int
	for x > 0 {
		// Add the current digit into result
		result = result*10 + x%10

		// Check if the result is within MaxInt32 and MinInt32 bounds
		if signMultiplier*result >= math.MaxInt32 || signMultiplier*result <= math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	// Restore to original sign of number (+ or -)
	return signMultiplier * result
}

func reverse(i int) int {
    ret := 0
    pos := 0
    if(i > 0) {
        pos = 1
    }
    if(pos == 0) {
        i = -1 * i
    }

    var iArr []string
    j := 10
    for i % j != i {
        d := (i % j) / (j / 10)
        i = i - d
        iArr = append(iArr, strconv.Itoa(d))
        j = j * 10
    }
    iArr = append(iArr, strconv.Itoa(i/(j/10)))

    ret, _ = strconv.Atoi(strings.Join(iArr, ""))
    if(pos == 0) {
        ret = -1 * ret
    }
    if(ret <= math.MinInt32 || ret >= math.MaxInt32) {
        ret = 0
    }
    return ret
}

func main() {
    fmt.Println("100: ", reverse(100))
    fmt.Println("1100: ", reverse(1100))
    fmt.Println("123: ", reverse(123))
    fmt.Println("-123: ", reverse(-123))
    fmt.Println("0: ", reversePro(0))
    fmt.Println("0: ", reverse(0))
}
