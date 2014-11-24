package main

/**
 *
 * Coded triangle numbers
 *
 * Problem 42
 *
 * The nth term of the sequence of triangle numbers is given by, t(n) = 0.5n(n+1); so the first ten triangle numbers are:
 *
 * 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
 *
 * By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value.
 * For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.
 * 
 * Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
 *
 */
 
import (
    "fmt"
    "time"
    "encoding/csv"
    "io"
    "os"
)


func contains(s []float64, e float64) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func main() {
    start := time.Now()
    
    tri := []float64{}
    
    for i := 1; i < 100; i++ {
        tri = append(tri, 0.5 * float64(i) * (float64(i)+1))
    }

    score := map[string]int{
        "A": 1,
        "B": 2,
        "C": 3,
        "D": 4,
        "E": 5,
        "F": 6,
        "G": 7,
        "H": 8,
        "I": 9,
        "J": 10,
        "K": 11,
        "L": 12,
        "M": 13,
        "N": 14,
        "O": 15,
        "P": 16,
        "Q": 17,
        "R": 18,
        "S": 19,
        "T": 20,
        "U": 21,
        "V": 22,
        "W": 23,
        "X": 24,
        "Y": 25,
        "Z": 26,
    }

    output := 0
    
    file, err := os.Open("p042_words.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    reader := csv.NewReader(file)
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
            return
        }
        
        fmt.Println(record) // record has the type []string
        
        for _, word := range record {
            total := 0
            
            for _, r := range word {
                c := string(r)
                total += score[c]
            }
            
            if contains(tri, float64(total)) {
                output++
            }	
        }

    }
    
    fmt.Println("Answer: ", output);

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}
