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
     
    var str string
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter an array of elements (int/string) in a line ")
    fmt.Println("")
    fmt.Println( "For eg: str1 str2 34 43 str3 ")
    fmt.Println("")
    scanner.Scan()
    input:= scanner.Text()
    
    fmt.Println("")
    fmt.Println("Enter any string OR integer to search : ")
    fmt.Println("")
    fmt.Scan(&str)

     var strArray []string= strings.Fields(input)
     
     fmt.Println("")
     fmt.Print("member found : ", IsMember(str,strArray))

}


func IsMember(str string, strArray []string ) bool{
	
   for _, value := range strArray {
      if value == str {
         return true
      }
   }
   return false
}
