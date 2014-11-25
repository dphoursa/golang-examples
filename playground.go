package main

import (
    "fmt"
    "time"
    "strconv"
    "sort"
    "strings"
    //"math"
)


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
    
    fmt.Println(1193, x[1193], x[1931], x[9311], x[3119]);
    
    
    
    for i := 0; i < threshold; i++ {
        valid := true
        
        if (x[i] == 0) {
            continue;
        }
        
        perm := i
        
        if( i == 1193 ) {
            fmt.Println("START:", i);
        }
        
        //find all ints in string distinct fact if one add else permutate

        temp := SortString(strconv.Itoa(perm))
        
        if strings.Contains(temp, "0") {
            continue
        }
        
        perm, _ = strconv.Atoi(temp)
        
        for true {
            if( i == 1193 ) {
                fmt.Println("IS PRIME:", perm, x[perm]);
            }
            if x[perm] == 0 {
                valid = false
                break
            }
        
            perm = permutate(perm);
            if perm == 0 {
                break;
            }

        }
        
        if valid {
            output = append(output, i)
        }
        
        if( i == 1193 ) {
            fmt.Println("END:", i);
        }
    }
    
    for i := 0; i < len(output); i++ {
        fmt.Println(output[i]);
    }
    
    fmt.Println("Answer:", len(output));

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func fact(n int) int {
  if n == 0 {
    return 1
  } else {
    return n * fact(n - 1)
  }
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
        if tmp < 0 || numbers[tmp] < numbers[i] {
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
    
    number, _ = strconv.Atoi(start + end)
    
    //4. Reverse the sequence from a[k + 1] up to and including the final element a[n].
    return number;
}
