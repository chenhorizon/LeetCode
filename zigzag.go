package main

import "fmt"
import "strings"

func convert(s string, numRows int) string {

    if(numRows == 1) {
        return s
    }

    sArr := strings.Split(s, "")
    sMap := make(map[int][]string)
    
    iRow := 0
    iIndex := 1

    for i:=0; i<len(sArr); i++ {
        if(i == 0) {
            iRow = 0
            sMap[iRow] = append(sMap[iRow], sArr[i])
            continue
        }
        if(iIndex == -1) {
            iRow = iRow - 1
            sMap[iRow] = append(sMap[iRow], sArr[i])
        }else if(iIndex == 1) {
            iRow = iRow + 1
            sMap[iRow] = append(sMap[iRow], sArr[i])
        }
        if(i % (numRows-1) == 0) {
            iIndex = -1 * iIndex
        }
    }

    ret := ""
    for i:=0; i<len(sMap); i++ {
        ret = ret + strings.Join(sMap[i], "")
    }

    return ret
}

//    0  6    12
//    1 57  1113
//    24 810  14
//    3  9    15

func main() {
    fmt.Println(convert("PAYPALISHIRING", 3))
    fmt.Println(convert("PAYPALISHIRING", 4))
    fmt.Println(convert("AB", 1))
}
