
package main 

import (

   "fmt"
   "strings"
  )

func main() {

	var input string
	fmt.Println("Enter your input string:")
	fmt.Scanf("%s", &input)
  input = strings.ToLower(input)
	fmt.Println(IsPalindrome(input))
}


func IsPalindrome(input string) string  {
  var result string

// reverse the input string and compare the result with input string  
  //here, index not required. So, using blank identifier _.

  for _,v := range input {
    result= string(v) + result
  }
 	if input== result {
  		return input+ " is a Palindrome String."
 	}
 
   return input+ " is NOT a Palindrome String."
}


