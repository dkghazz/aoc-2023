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
        isPartNumber := false
        number := 0
        for j, c := range line {
            if c >= ‘0’ && c <= ‘9’ {
                number = number*10 + int(c-‘0’)
                if j+1 <= len(line)-1 && isSymbol(mapp[i][j+1]) || j-1 >= 0 && isSymbol(mapp[i][j-1]) || i+1 <= len(mapp)-1 && isSymbol(mapp[i+1][j]) || i-1 >= 0 && isSymbol(mapp[i-1][j]) ||
                    i+1 <= len(mapp)-1 && j+1 <= len(line)-1 && isSymbol(mapp[i+1][j+1]) || i-1 >= 0 && j-1 >= 0 && isSymbol(mapp[i-1][j-1]) || i+1 <= len(mapp)-1 && j-1 >= 0 && isSymbol(mapp[i+1][j-1]) || i-1 >= 0 && j+1 <= len(line)-1 && isSymbol(mapp[i-1][j+1]) {
                    isPartNumber = true
                }
            } else if !isNumber(c) && number != 0 && isPartNumber {
                answer += number
                number = 0
                isPartNumber = false
            } else {
                number = 0
                isPartNumber = false
            }
            // println(number)
        }
        if number != 0 && isPartNumber {
            answer += number
        }
    }
    println(answer)
}