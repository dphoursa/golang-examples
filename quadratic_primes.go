package main

import (
    "fmt"
    "time"
)

/**
 *
 * Quadratic primes
 *
 * Problem 27
 * 
 * Euler discovered the remarkable quadratic formula: n^2 + n + 41
 *
 * FIt turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39.
 * However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.
 * 
 * The incredible formula  n^2 - 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79.
 * The product of the coefficients, -79 and 1601, is -126479.
 *
 * Considering quadratics of the form:
 *
 * n^2 + an + b, where |a| < 1000 and |b| < 1000
 *
 * where |n| is the modulus/absolute value of n
 * e.g. |11| = 11 and |-4| = 4
 *
 * Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
 *
 */

func main() {
    start := time.Now()
    max_threshold := 10000000

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
    output_a := 0
    output_b := 0
    
    // n^2 + an + b, where |a| < 1000 and |b| < 1000
    
    for a := 999; a >= -999; a-- {
        tmp_output := []int{}
        tmp_output_b := 0
        for b := 999; b >= -999; b-- {
            tmp_output_2 := []int{}
            for n := 0; n < 100000; n++ {
                answer := n * n + a * n + b
                if answer < 0 || x[answer] == 0 {
                    break
                }
                
                tmp_output_2 = append(tmp_output_2, answer)
            }
            
            if len(tmp_output_2) >= len(tmp_output) {
                tmp_output_b = b
                tmp_output = tmp_output_2
            }
        }
        
        if len(tmp_output) >= len(output) {
            output_a = a
            output_b = tmp_output_b
            output = tmp_output
        }
    }
       
    

    fmt.Println("Answer:", output_a * output_b);

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}