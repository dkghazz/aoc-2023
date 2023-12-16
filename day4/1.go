package main
import (
    “bufio”
    “math”
    “os”
    “strings”
)
func main() {
    answer := 0
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if line == “” {
            break
        }
        barSeparated := strings.Split(strings.Split(line, “:”)[1], “|”)
        winnings := strings.Split(barSeparated[0], ” “)
        numbers := strings.Split(barSeparated[1], ” “)
        winningCount := 0
        for _, number := range numbers {
            if number == “” {
                continue
            }
            for _, winning := range winnings {
                if winning == “” {
                    continue
                }
                if number == winning {
                    winningCount++
                    break
                }
            }
        }
        answer += int(math.Pow(2, float64(winningCount-1)))
    }
    println(answer)
}