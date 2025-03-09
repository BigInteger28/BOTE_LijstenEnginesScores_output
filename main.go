package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// getRating functie blijft ongewijzigd
func getRating(maxPos int, score int) (int, int) {
    var level int
    var rating int
    var ps31 int = (maxPos * 310) / 1000
    var ps50 int = (maxPos * 494) / 1000
    var ps60 int = (maxPos * 592) / 1000
    if (score * 100) / maxPos < 50 {
        rating = ((score - ps31) * 5) + 1095
    } else if (score * 100) / maxPos < 60 {
        rating = ((score - ps50) * 10) + 2500
    } else {
        rating = ((score - ps60) * 20) + 4050
    }
    if rating < 500 {
        return 1, rating
    }
    level = ((rating - 500) / 75) + 2
    return level, rating
}

func main() {
    var maxPos int
    var names []string
    var scores []int
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Max punten mogelijk: ")
    if scanner.Scan() {
        maxPosStr := scanner.Text()
        var err error
        maxPos, err = strconv.Atoi(maxPosStr)
        if err != nil {
            fmt.Println("Ongeldige maxPos, probeer opnieuw.")
            return
        }
    }

    fmt.Println("Voer namen in, één per lijn. Voer '.' in om te stoppen.")
    for scanner.Scan() {
        name := scanner.Text()
        if name == "." {
            break
        }
        names = append(names, name)
    }

    fmt.Println("Voer scores in, één per lijn. Voer '.' in om te stoppen.")
    for scanner.Scan() {
        line := scanner.Text()
        if line == "." {
            break
        }
        score, err := strconv.Atoi(line)
        if err != nil {
            fmt.Println("Ongeldige score, probeer opnieuw.")
            continue
        }
        scores = append(scores, score)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Fout bij het lezen van invoer:", err)
        return
    }

    if len(names) != len(scores) {
        fmt.Println("Aantal namen en scores komen niet overeen.")
        return
    }

    // OUTPUT
    fmt.Println("NAAM   LEVEL   RATING")
    for i := 0; i < len(names); i++ {
        level, rating := getRating(maxPos, scores[i])
        fmt.Printf("%s   %d   %d\n", names[i], level, rating)
    }
}
