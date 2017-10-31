package main 

import (

   "fmt"
   "strings"
   "bufio"
    "os"

)

func main() {
       scanner := bufio.NewScanner(os.Stdin)
        var input string
        fmt.Print("Enter your input string: ")
        scanner.Scan()
        input = scanner.Text()
        fmt.Println("Corrected string is:",SpellCorrect(input))
        
    }


func SpellCorrect(input string) string {
	
 // add space if there is no space after dot (.) AND 
 // divide the input into fields and add space after each field
 
	input= strings.Replace(input, ".", ". ", -1)

 	result:= strings.Join(strings.Fields(input), " ")

	return result
}


