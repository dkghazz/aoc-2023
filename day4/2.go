package main
import (
    “bufio”
    “os”
    “strings”
)
func main() {
    answer := 0
    scanner := bufio.NewScanner(os.Stdin)
    idx := 0
    var cards []int
    for scanner.Scan() {
        line := scanner.Text()
        if line == “” {
            break
        }
        if len(cards)-1 < idx {
            cards = append(cards, 1)
        } else {
            cards[idx]++
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
        for i := 1; i <= winningCount; i++ {
            if len(cards)-1 < idx+i {
                cards = append(cards, cards[idx])
            } else {
                cards[idx+i] += cards[idx]
            }
        }
        idx++
    }
    for _, card := range cards {
        answer += card
    }
    println(answer)
}