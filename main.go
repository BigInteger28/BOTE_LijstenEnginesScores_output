package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// getRating functie om level en rating te berekenen
func getRating(maxPos int, score int) (int, int) {
	var level, rating, perM int
	//      5.26    <25   30   40    45    50    53    55    57    60    65
	var PRM = []int {0, 25, 925, 1275, 2500, 3950, 4700, 5300, 5750, 6275}
	//               0  131  73    263   276   164   115   89    29     0  
	//var PRM = []int {250, 500, 1175, 1625, 2975, 3875, 4325, 4775, 5675, 6050}
	perM = (score * 1000) / maxPos
	
	if perM < 250 {
		rating = ((PRM[0] / (25-0)) * perM) / 10 // 2 rating per promile onder 250 promile
	} else if perM < 300 {
		rating = PRM[0] + ((((PRM[1]-PRM[0]) / (30-25)) * (perM-250)) / 10) //X rating per promile boven 250 promile 
	} else if perM < 400 {
		rating = PRM[1] + ((((PRM[2]-PRM[1]) / (40-30)) * (perM-300)) / 10) //X rating per promile boven 300 promile
	} else if perM < 450 {
		rating = PRM[2] + ((((PRM[3]-PRM[2]) / (45-40)) * (perM-400)) / 10) //X rating per promile boven 300 promile
	} else if perM < 500 {
		rating = PRM[3] + ((((PRM[4]-PRM[3]) / (50-45)) * (perM-450)) / 10) //X rating per promile boven 400 promile
	} else if perM < 530 {
		rating = PRM[4] + ((((PRM[5]-PRM[4]) / (53-50)) * (perM-500)) / 10) //X rating per promile boven 500 promile
	} else if perM < 550 {
		rating = PRM[5] + ((((PRM[6]-PRM[5]) / (55-53)) * (perM-530)) / 10) //X rating per promile boven 530 promile
	} else if perM < 570 {
		rating = PRM[6] + ((((PRM[7]-PRM[6]) / (57-55)) * (perM-550)) / 10) //X rating per promile boven 550 promile
	} else if perM < 600 {
		rating = PRM[7] + ((((PRM[8]-PRM[7]) / (60-57)) * (perM-570)) / 10) //X rating per promile boven 570 promile
	} else if perM < 650 {
		rating = PRM[8] + ((((PRM[9]-PRM[8]) / (65-60)) * (perM-600)) / 10) //12 rating per promile ONDER 650 PROMILE
	} else {
		rating = PRM[9] + ((perM-700)*20) //20 rating per promile boven 700 promile
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

    // Scanner om invoer te lezen
    scanner := bufio.NewScanner(os.Stdin)

    // Max punten mogelijk invoeren
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

    // Namen invoeren
    fmt.Println("Voer namen in, één per lijn. Voer '.' in om te stoppen.")
    for scanner.Scan() {
        name := scanner.Text()
        if name == "." {
            break
        }
        names = append(names, name)
    }

    // Scores invoeren
    fmt.Println("Voer scores in, één per lijn. Voer '.' in om te stoppen.")
    for scanner.Scan() {
        line := scanner.Text()
        if line == "." {
            break
        }
        // Splits de invoer op "/"
        parts := strings.Split(line, "/")
        if len(parts) < 1 {
            fmt.Println("Ongeldige score, probeer opnieuw.")
            continue
        }
        // Neem het eerste deel en verwijder spaties
        scoreStr := strings.TrimSpace(parts[0])
        // Zet om naar integer
        score, err := strconv.Atoi(scoreStr)
        if err != nil {
            fmt.Println("Ongeldige score, probeer opnieuw.")
            continue
        }
        scores = append(scores, score)
    }

    // Controleer op fouten in de scanner
    if err := scanner.Err(); err != nil {
        fmt.Println("Fout bij het lezen van invoer:", err)
        return
    }

    // Controleer of het aantal namen en scores overeenkomt
    if len(names) != len(scores) {
        fmt.Println("Aantal namen en scores komen niet overeen.")
        return
    }

    // Resultaten weergeven
    fmt.Println("NAAM   LEVEL   RATING")
    for i := 0; i < len(names); i++ {
        level, rating := getRating(maxPos, scores[i])
        fmt.Printf("%s   %d   %d\n", names[i], level, rating)
    }
	fmt.Scanln()
}
