package main

/**
 *
 * Reciprocal cycles
 * 
 * Problem 26
 * 
 * A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:
 * 
 *     1/2	= 	0.5
 *     1/3	= 	0.(3)
 *     1/4	= 	0.25
 *     1/5	= 	0.2
 *     1/6	= 	0.1(6)
 *     1/7	= 	0.(142857)
 *     1/8	= 	0.125
 *     1/9	= 	0.(1)
 *     1/10	= 	0.1
 * 
 * Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.
 *
 * Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
 *
 */
import (
    "strconv"
    "fmt"
    "time"
)

func main() {
    start := time.Now()

    fmt.Println("Answer: ", compute(1000));	

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}

func contains(s []int, e int) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func compute(threshold int) string {
    output := 0
    cycles := 0
    for i := 2; i < threshold; i++ {
        count := 0
        noms := []int{}
        nom := 1
        for true {
            nom = (nom * 10) % i
            if nom == 0 {
                break
            } else if contains(noms, nom) {
                if cycles < count {
                    cycles = count
                    output = i
                }
                break
            }
            count++
            noms = append(noms, nom)
        }
    }
    return strconv.Itoa(output)
}

