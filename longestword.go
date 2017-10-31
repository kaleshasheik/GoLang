/*** author 
     BGH25333
 ***/
     

package main

import (
 "fmt"
 "bufio"
 "os"
 "strings"
 
)

func main(){
   
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter an array of strings in a line ")
    fmt.Println("")
    fmt.Println( "For eg: str1 str2 str3 ")
    fmt.Println("")
    scanner.Scan()
    input:= scanner.Text()
    var strArray []string= strings.Fields(input)
    fmt.Println("")
    fmt.Println("")
    FindLongestWord(strArray)
    

    }

    func FindLongestWord(strArray []string)  {
    	
        max := strArray[0] 

         for _, value := range strArray {
         
                 if len(value) > len(max) {
                         max = value 
                 }
         }
     
     fmt.Println("Longest string is :",max, " and length is :" ,len(max))

    }