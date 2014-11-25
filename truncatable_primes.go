package main

import (
    "fmt"
    "time"
    "strconv"
    "strings"
)

/**
 *
 * Truncatable primes
 *
 * Problem 37
 * 
 * The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and
 * remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
 *
 * Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
 * 
 * NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
 *
 */

func main() {
    start := time.Now()
    max_threshold := 1000000

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
    
    for i := 10; i < max_threshold; i++ {
        valid := true
        
        if (x[i] == 0) {
            continue;
        }
        
        temp := strconv.Itoa(i)
        
        if strings.Contains(temp, "0") {
            continue
        }
        
        for i := 1; i < len(temp); i++ {
           
            var temp_2 = temp[0:len(temp) - i]
            var perm, _ = strconv.Atoi(temp[i:])
            var perm_2, _ = strconv.Atoi(temp_2)

            if x[perm] == 0 || x[perm_2] == 0 {
                valid = false
                break
            }
        }
        
        if valid {
            output = append(output, i)
        }

    }
    
    sum := 0
    for _, value := range output {
        sum += value
        fmt.Println(value);
    }

    fmt.Println("Answer:", sum);

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}