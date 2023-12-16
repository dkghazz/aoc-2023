package main
import (
    “bufio”
    “fmt”
    “os”
    “strconv”
)
func main() {
    answer := 0
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if line == “” {
            break
        }
        var firstDigit byte
        for i := 0; i < len(line); i++ {
            if i+2 < len(line) && line[i] == ‘o’ && line[i+1] == ‘n’ && line[i+2] == ‘e’ {
                firstDigit = ‘1’
                break
            } else if i+2 < len(line) && line[i] == ‘t’ && line[i+1] == ‘w’ && line[i+2] == ‘o’ {
                firstDigit = ‘2’
                break
            } else if i+3 < len(line) && line[i] == ‘t’ && line[i+1] == ‘h’ && line[i+2] == ‘r’ && line[i+3] == ‘e’ {
                firstDigit = ‘3’
                break
            } else if i+3 < len(line) && line[i] == ‘f’ && line[i+1] == ‘o’ && line[i+2] == ‘u’ && line[i+3] == ‘r’ {
                firstDigit = ‘4’
                break
            } else if i+3 < len(line) && line[i] == ‘f’ && line[i+1] == ‘i’ && line[i+2] == ‘v’ && line[i+3] == ‘e’ {
                firstDigit = ‘5’
                break
            } else if i+2 < len(line) && line[i] == ‘s’ && line[i+1] == ‘i’ && line[i+2] == ‘x’ {
                firstDigit = ‘6’
                break
            } else if i+4 < len(line) && line[i] == ‘s’ && line[i+1] == ‘e’ && line[i+2] == ‘v’ && line[i+3] == ‘e’ && line[i+4] == ‘n’ {
                firstDigit = ‘7’
                break
            } else if i+4 < len(line) && line[i] == ‘e’ && line[i+1] == ‘i’ && line[i+2] == ‘g’ && line[i+3] == ‘h’ && line[i+4] == ‘t’ {
                firstDigit = ‘8’
                break
            } else if i+3 < len(line) && line[i] == ‘n’ && line[i+1] == ‘i’ && line[i+2] == ‘n’ && line[i+3] == ‘e’ {
                firstDigit = ‘9’
                break
            } else if line[i] >= ‘1’ && line[i] <= ‘9’ {
                firstDigit = line[i]
                break
            }
        }
        var lastDigit byte
        for i := len(line) - 1; i >= 0; i-- {
            if i+2 < len(line) && line[i] == ‘o’ && line[i+1] == ‘n’ && line[i+2] == ‘e’ {
                lastDigit = ‘1’
                break
            } else if i+2 < len(line) && line[i] == ‘t’ && line[i+1] == ‘w’ && line[i+2] == ‘o’ {
                lastDigit = ‘2’
                break
            } else if i+3 < len(line) && line[i] == ‘t’ && line[i+1] == ‘h’ && line[i+2] == ‘r’ && line[i+3] == ‘e’ {
                lastDigit = ‘3’
                break
            } else if i+3 < len(line) && line[i] == ‘f’ && line[i+1] == ‘o’ && line[i+2] == ‘u’ && line[i+3] == ‘r’ {
                lastDigit = ‘4’
                break
            } else if i+3 < len(line) && line[i] == ‘f’ && line[i+1] == ‘i’ && line[i+2] == ‘v’ && line[i+3] == ‘e’ {
                lastDigit = ‘5’
                break
            } else if i+2 < len(line) && line[i] == ‘s’ && line[i+1] == ‘i’ && line[i+2] == ‘x’ {
                lastDigit = ‘6’
                break
            } else if i+4 < len(line) && line[i] == ‘s’ && line[i+1] == ‘e’ && line[i+2] == ‘v’ && line[i+3] == ‘e’ && line[i+4] == ‘n’ {
                lastDigit = ‘7’
                break
            } else if i+4 < len(line) && line[i] == ‘e’ && line[i+1] == ‘i’ && line[i+2] == ‘g’ && line[i+3] == ‘h’ && line[i+4] == ‘t’ {
                lastDigit = ‘8’
                break
            } else if i+3 < len(line) && line[i] == ‘n’ && line[i+1] == ‘i’ && line[i+2] == ‘n’ && line[i+3] == ‘e’ {
                lastDigit = ‘9’
                break
            } else if line[i] >= ‘1’ && line[i] <= ‘9’ {
                lastDigit = line[i]
                break
            }
        }
        lineAnswer, _ := strconv.Atoi(fmt.Sprintf(“%x%x”, firstDigit-‘0’, lastDigit-‘0’))
        answer += lineAnswer
    }
    println(answer)
}