package main
import (
    “bufio”
    “os”
)
func isSymbol(c byte) bool {
    return !isNumber(c) && c != ‘.’
}
func isNumber(c byte) bool {
    return c >= ‘0’ && c <= ‘9’
}
func main() {
    answer := 0
    scanner := bufio.NewScanner(os.Stdin)
    var mapp [][]byte
    for scanner.Scan() {
        line := scanner.Text()
        if line == “” {
            break
        }
        mapp = append(mapp, []byte(line))
    }
    for i, line := range mapp {
        for j, c := range line {
            findNumber := func(r, c int) int {
                startC := c
                endC := c
                for i := 1; ; i++ {
                    if c-i < 0 || !isNumber(mapp[r][c-i]) {
                        break
                    }
                    startC = c - i
                }
                for i := 1; ; i++ {
                    if c+i >= len(mapp[r]) || !isNumber(mapp[r][c+i]) {
                        break
                    }
                    endC = c + i
                }
                var ret int
                for i := startC; i <= endC; i++ {
                    ret = ret*10 + int(mapp[r][i]-‘0’)
                }
                return ret
            }
            adjacentNumbers := func(r, c int) (numOne, numTwo int) {
                if c-1 >= 0 && isNumber(mapp[r][c-1]) {
                    numOne = findNumber(r, c-1)
                }
                if c+1 < len(mapp[r]) && isNumber(mapp[r][c+1]) {
                    if numOne == 0 {
                        numOne = findNumber(r, c+1)
                    } else {
                        numTwo = findNumber(r, c+1)
                    }
                }
                if r-1 >= 0 && isNumber(mapp[r-1][c]) {
                    if numOne == 0 {
                        numOne = findNumber(r-1, c)
                    } else if numTwo == 0 {
                        numTwo = findNumber(r-1, c)
                    } else {
                        return 0, 0
                    }
                } else {
                    if r-1 >= 0 && c-1 >= 0 && isNumber(mapp[r-1][c-1]) {
                        if numOne == 0 {
                            numOne = findNumber(r-1, c-1)
                        } else if numTwo == 0 {
                            numTwo = findNumber(r-1, c-1)
                        } else {
                            return 0, 0
                        }
                    }
                    if r-1 >= 0 && c+1 < len(mapp[r]) && isNumber(mapp[r-1][c+1]) {
                        if numOne == 0 {
                            numOne = findNumber(r-1, c+1)
                        } else if numTwo == 0 {
                            numTwo = findNumber(r-1, c+1)
                        } else {
                            return 0, 0
                        }
                    }
                }
                if r+1 < len(mapp) && isNumber(mapp[r+1][c]) {
                    if numOne == 0 {
                        numOne = findNumber(r+1, c)
                    } else if numTwo == 0 {
                        numTwo = findNumber(r+1, c)
                    } else {
                        return 0, 0
                    }
                } else {
                    if r+1 < len(mapp) && c-1 >= 0 && isNumber(mapp[r+1][c-1]) {
                        if numOne == 0 {
                            numOne = findNumber(r+1, c-1)
                        } else if numTwo == 0 {
                            numTwo = findNumber(r+1, c-1)
                        } else {
                            return 0, 0
                        }
                    }
                    if r+1 < len(mapp) && c+1 < len(mapp[r]) && isNumber(mapp[r+1][c+1]) {
                        if numOne == 0 {
                            numOne = findNumber(r+1, c+1)
                        } else if numTwo == 0 {
                            numTwo = findNumber(r+1, c+1)
                        } else {
                            return 0, 0
                        }
                    }
                }
                return numOne, numTwo
            }
            if c == ‘*’ {
                adjaOne, adjaTwo := adjacentNumbers(i, j)
                if adjaOne != 0 && adjaTwo != 0 {
                    answer += adjaOne * adjaTwo
                }
            }
        }
    }
    println(answer)
}