package main

/**
 *
 * Champernowne's constant
 *
 * Problem 40
 *
 * An irrational decimal fraction is created by concatenating the positive integers:
 *
 * 0.123456789101112131415161718192021...
 *
 * It can be seen that the 12th digit of the fractional part is 1.
 * 
 * If dn represents the nth digit of the fractional part, find the value of the following expression.
 *
 * d1 x d10 x d100 x d1000 x d10000 x d100000 x d1000000
 *
 */
 
import (
    "strconv"
    "fmt"
    "time"
)

func main() {
    start := time.Now()

    fmt.Println("Answer: ", compute());	

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}

func compute ()  int {
    output := "1"
    count := 1
    for len(output) <= 1000000 {
        count++
        output += strconv.Itoa(count)
    }
    
    i, _ := strconv.Atoi(output[0:1])
    i2, _ := strconv.Atoi(output[9:10])
    i3, _ := strconv.Atoi(output[99:100])
    i4, _ := strconv.Atoi(output[999:1000])
    i5, _ := strconv.Atoi(output[9999:10000])
    i6, _ := strconv.Atoi(output[99999:100000])
    i7, _ := strconv.Atoi(output[999999:1000000])
    
    fmt.Println(i, i2, i3, i4, i5, i6, i7);	
    
    return i * i2 * i3 * i4 * i5 * i6 * i7
}