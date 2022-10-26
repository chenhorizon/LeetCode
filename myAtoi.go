package main

import "fmt"
import "strings"
import "strconv"

func myAtoi(s string) int {
    ret := 0
    positive := true
    inNum := false

    s = strings.Trim(s, " ")
    sArr := strings.Split(s, "")

    for _, item := range(sArr) {

        switch item {
            case "-":
                inNum = true
                fmt.Println("-")
            case "+":
                inNum = true
                fmt.Println("+")
            case strconv.Atoi(item) >=0 || strconv.Atoi(item) <= 9:
                inNum = true
                fmt.Println(item)
        }
    }
    
    return ret
}

func main() {
    myAtoi("+123")
    myAtoi("-123")
    myAtoi("-1+3")
}
