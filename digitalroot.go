package main 

import (

   "fmt"
   
)

func main() {

	
var number int64
 
    fmt.Print("Enter any positive integer : ")
    fmt.Scan(&number)
    result:= FindSumOfDigits(number)
    
    fmt.Print("sum of digits is : ",result)
}

func FindSumOfDigits(number int64) int64 {
   
   var sum int64=0
   
   for number>0 {
   
   sum= sum+number%10
   number= number/10
 
  }
  // check if the sum is single digit
    for sum >9 {

        sum= FindSumOfDigits(sum)
    }

 return sum

}