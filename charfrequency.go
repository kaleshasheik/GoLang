/*** author 
     BGH25333
 ***/
     
package main

import (

   "fmt"
   "strings"
  )

func main() {
 
	var input string
	fmt.Print("Enter your input string: ")
  fmt.Println("For eg: abbabcbdbabdbdbabababcbcbab")
  fmt.Println("")
	fmt.Scanf("%s", &input)
  fmt.Println("")

	CharFrequency(input)
}

func CharFrequency(str string) {

   myMap:= make(map[string]int)
   var count int
   for _, v := range str {
     
    count = strings.Count(str, string(v))

  //check if key is already present in map. If not, then add into the map
   _, ok := myMap[string(v)]

   if ok == false {
         myMap[string(v)]= count
        
    } 
   }

    for k,v := range myMap {

       fmt.Println("Input char is :", k ," occurence :" ,v)
    }
   
   }
