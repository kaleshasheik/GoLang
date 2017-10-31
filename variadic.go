package main

import (
 "fmt"
 "bufio"
 "os"
 "strings"
)

func main(){
    
    scanner := bufio.NewScanner(os.Stdin)
    var input string
    fmt.Print("Enter the strings in a line : ")
    scanner.Scan()
    input = scanner.Text()
    var strArray [] string=strings.Fields(input)
    fmt.Println("concatenated string is: ",Concatenate(strArray...))
}

func Concatenate(input ...string) string {
    var result string

    // iterate and append the strings
  //here, index not required. So, using blank identifier _.
  
    for _,value := range input {
    result = result +value
  }
    return result
}