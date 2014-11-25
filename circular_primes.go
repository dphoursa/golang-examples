package main

import (
    "fmt"
    "time"
    "strconv"
    "strings"
)

/**
 *
 * Circular primes
 *
 * Problem 35
 * 
 * The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
 *
 * There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
 * 
 * How many circular primes are there below one million?
 *
 */

func main() {
    start := time.Now()
    max_threshold := 1000000
    threshold := 1000000
    

    x := make ([]int, max_threshold)
    
    for i := 0; i < max_threshold; i++ {
        x[i] = 1;
    }
    
    x[0] = 0
    x[1] = 0

    i := 2;
        
    for i < max_threshold {
        for j := 2; (j * i) < (max_threshold); j++ {
            tmp := j * i;
            x[tmp] = 0
        }
        
        i++;
        for i < max_threshold && x[i] == 0 {
            i++;
        }
    }
    
    output := []int{}
    
    for i := 0; i < threshold; i++ {
        valid := true
        
        if (x[i] == 0) {
            continue;
        }
        
        temp := strconv.Itoa(i)
        
        if strings.Contains(temp, "0") {
            continue
        }
        
        for i := 0; i < len(temp); i++ {
           
            temp = temp[1:] + temp[:1]
            var perm, _ = strconv.Atoi(temp)
           
            if x[perm] == 0 {
                valid = false
                break
            }
        }
        
        if valid {
            output = append(output, i)
        }

    }
    
    for i := 0; i < len(output); i++ {
        fmt.Println(output[i]);
    }
    
    fmt.Println("Answer:", len(output));

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}