package main
import (
    “bufio”
    “os”
    “strconv”
    “strings”
)
func parseInputLine(line string) []string {
    if line[0] != ‘G’ && line[1] != ‘a’ && line[2] != ‘m’ && line[3] != ‘e’ {
        return nil
    }
    prefixRemoved := strings.Split(line, “:”)[1]
    parsed := strings.Split(prefixRemoved, “;”)
    return parsed
}
func main() {
    answer := 0
    scanner := bufio.NewScanner(os.Stdin)
    // game := 1
    for scanner.Scan() {
        line := scanner.Text()
        if line == “” {
            break
        }
        parsed := parseInputLine(line)
        if parsed == nil {
            break
        }
        // impossible := false
        r := 0
        g := 0
        b := 0
        for i := 0; i < len(parsed); i++ {
            parsedRGB := strings.Split(parsed[i], “,”)
            for j := 0; j < len(parsedRGB); j++ {
                for k := 0; k < len(parsedRGB[j]); k++ {
                    if parsedRGB[j][k] == ‘r’ {
                        num, _ := strconv.Atoi(parsedRGB[j][1 : k-1])
                        if r < num {
                            r = num
                        }
                        // if num > 12 {
                        //  impossible = true
                        // }
                        break
                    }
                    if parsedRGB[j][k] == ‘g’ {
                        num, _ := strconv.Atoi(parsedRGB[j][1 : k-1])
                        if g < num {
                            g = num
                        }
                        // if num > 13 {
                        //  impossible = true
                        // }
                        break
                    }
                    if parsedRGB[j][k] == ‘b’ {
                        num, _ := strconv.Atoi(parsedRGB[j][1 : k-1])
                        if b < num {
                            b = num
                        }
                        // if num > 14 {
                        //  impossible = true
                        // }
                        break
                    }
                }
            }
        }
        // if !impossible {
        //  answer += game
        // }
        // game++
        answer += r * g * b
    }
    println(answer)
}