package main

import (
    "fmt"
    "time"
    "strconv"
    //"math"
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
    threshold := 1000000

    x := make ([]int, threshold)
    
    for i := 0; i < threshold; i++ {
        x[i] = 1;
    }
    
    x[0] = 0
    x[1] = 0
    
    var primes []int
    
    i := 2;
        
    for i < threshold {
        primes = append(primes, i)
        for j := 2; (j * i) < (threshold); j++ {
            tmp := j * i;
            x[tmp] = 0
        }
        
        i++;
        for i < threshold && x[i] == 0 {
            i++;
        }
    }
    
    output := []int{}
    
    for i := 0; i < len(primes); i++ {
        valid := true
        perm := primes[i]
        for true {
            perm := permutate(perm);
            
            if perm != 0 {
                break;
            }
            
            if x[perm] == 0 {
                valid = false
                break
            }
        }
        
        if valid {
            output = append(output, primes[i])
        }
    }
    
    for i := 0; i < len(output); i++ {
        fmt.Println(output[i]);
    }

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func permutate(number int) int {
    numbers := []rune(strconv.Itoa(number))
    
    //1. Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
    i := len(numbers) - 1
    
    //for isset(numbers[i - 1]) && numbers[i - 1] >= numbers[i] {
    
    for true {
        tmp := i - 1
        val, ok := numbers[tmp]
        if !ok || val < numbers[i] {
            break
        }
        i--
    }
    
    if i == 0 {
        return 0
    }

    //2. Find the largest index l greater than k such that a[k] < a[l].
    j := len(numbers) - 1;
    for numbers[j] <= numbers[i - 1] {
        j--;
    }
    
    //3. Swap the value of a[k] with that of a[l].
    tmp := numbers[i - 1];
    numbers[i - 1] = numbers[j];
    numbers[j] = tmp;
    
    start := string(numbers[:i])
    end := Reverse(string(numbers[i:]))
    
    //4. Reverse the sequence from a[k + 1] up to and including the final element a[n].
    return strconv.Atoi(start + end);
}
