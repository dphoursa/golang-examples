package main
 
import (
    "fmt"
    "time"
    "os"
)

func main() {
    start := time.Now()
    str := os.Args[1]

    permutate(str,0,len(str)) // call the function.

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}

// function to generate and print all N! permutations of $str. (N = strlen($str)).
func permutate(str string, i, n int) {
    if i == n {
        println(str)
    } else {
        for j := i; j < n; j++ {
            str = swap(str, i, j)
            permutate(str, i+1, n)
            str = swap(str, i, j) // backtrack.
        }
    }
}

// function to swap the char at pos $i and $j of $str.
func swap(str string, i, j int) string {
    temp := string(str[i])
    temp2 := string(str[j])
    str = str[:i] + temp2 + str[i+1:]
    str = str[:j] + temp + str[j+1:]
    return str
}

