package main

import (
    "fmt"
    
)

func main() {

   var number float64
 
    fmt.Print("Enter any number : ")
    fmt.Scan(&number)
    result := FindSquareRoot(float64(number))
    fmt.Println("square root is :" ,result)
     
    }

    func FindSquareRoot(input float64) float64 {
    if input == 0 { return 0 }
    oldApprox := 1.0
    var result float64

    //Newton law:  //    newApproximation =  (oldApproximation + N/oldApproximation)/2
    
    for i := 0; i < int(input); i++ {
            
        result= (oldApprox+input/oldApprox)/2
        oldApprox=result
    }
    return result
}    

