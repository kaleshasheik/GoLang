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
   
    fmt.Println("")
    arry1 := []string{"abc", "bcd", "43", "cde", "def","53","63","efg"}
    fmt.Println("Input array1 : ", arry1)
    fmt.Println("")
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter an array of elements (int/string) in a line ")
    fmt.Println("")
    fmt.Println( "For eg: str1 str2 34 43 str3 ")
    fmt.Println("")
    fmt.Print("Input array2 :" )
    fmt.Println("")
    scanner.Scan()
    input:= scanner.Text()
   
     var strArray []string= strings.Fields(input)
  
     fmt.Println("")
     fmt.Print("member found : ", SearchString(arry1,strArray))

}


func SearchString(strArray1 []string, strArray2 []string ) bool{
	   
    for _, arr1 := range strArray1 {

      for _, arr2 := range strArray2 {

      if arr1 == arr2 {
         return true
      }
    }
   }
   return false
}
