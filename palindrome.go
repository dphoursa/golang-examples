package main
 
import (
    "fmt";
    "time"
    "os"
    "strings"
)
 
func main() {
    if len(os.Args) < 2 {
        fmt.Print("Please provide a word as an input!")
	fmt.Print("\n")	
	os.Exit(1)
    }

    start := time.Now()
    str := os.Args[1]

    rev := Reverse(str)
    count := len(str)
    
    var result = ""

    for i := 0; i < count-1; i++ {
        for j := count; j > i; j-- {
            if strings.Contains(rev, str[i:j]) && len(string(str[i:j])) > len(string(result)) {
                result = str[i:j]
            }
        }
    }

    println(result)

    fmt.Print("Total time: ")
    fmt.Println(time.Since(start).Seconds())  
}

func Reverse(s string) (result string) {
    for _,v := range s {
        result = string(v) + result
    }
    return
}

