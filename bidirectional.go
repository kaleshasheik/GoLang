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
    longest := make(chan string)
    palindrome := make(chan []string)
    go FindLongestWord(longest,strArray)
    go CheckPalindrome(palindrome,strArray)
    
     fmt.Println("Received longest word from channel: ",<-longest)

      fmt.Println("Received palindrome strings from channel: ",<-palindrome)

    }

    func FindLongestWord(longest chan string, strArray []string)  {
        
        max := strArray[0] 

         for _, value := range strArray {
         
                 if len(value) > len(max) {
                         max = value 
                 }
         }
     
     longest<- max

    }

func CheckPalindrome(palindrome chan []string, strArray []string)  {

   
  var result []string

   for _,value := range strArray {

       isOK:= IsPalindrome(value)

       if(isOK ==true){

        result = append(result, value)
       }

   }

palindrome <- result

}


 func IsPalindrome(input string) bool  {

  var result string

// reverse the input string and compare the result with input string  
  //here, index not required. So, using blank identifier _.

  for _,v := range input {
    result= string(v) + result
  }
    if input== result {
        return true
    }
 
   return false
}

