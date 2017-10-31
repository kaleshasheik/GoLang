/*** author 
     BGH25333
 ***/
     

package main

import(
"fmt"
"regexp"
"sort"
 "bufio"
 "os"
)

type Item struct {
   
     Name       string
     Occurence  int
     SubItems []SubItem
}

type SubItem struct {
    
    Name string
    Occurence   int
}

func main() {
   
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter an array of strings in a line with delimeter")
    fmt.Println("")
    fmt.Println( "For eg: str1:str2,str3 ")
    fmt.Println("")
    scanner.Scan()
    input:= scanner.Text()

    re := regexp.MustCompile(`[A-Za-z']+`)
    tokens := re.FindAllString(input, -1)

    myMap:= StoreTokens(tokens)
    SortByKeyAndValue(myMap)

    }

func StoreTokens(str []string) map[string]int{

   myMap:= make(map[string]int)
   for _, v := range str {
     
  //check if key is already present in map. If not, then add elements into the map with count
   _, ok := myMap[string(v)]

   if ok == false {
         myMap[string(v)]= 1
        
    } else {
        myMap[string(v)] += 1
    }
   
   }
    return myMap
   }

 func SortByKeyAndValue(myMap map[string]int) {

  var item Item

  item.SubItems = make([]SubItem, len(myMap))
 
   k:=0
   for i,j := range myMap {

    item.SubItems[k] = SubItem{
        
        Name: i,
        Occurence:   j,
    }
    k++
}


  sort.Slice(item.SubItems, func(i, j int) bool { return item.SubItems[i].Occurence > item.SubItems[j].Occurence})
	fmt.Println("sorted by occurence in desending order :", item.SubItems)

	sort.Slice(item.SubItems, func(i, j int) bool { return item.SubItems[i].Name > item.SubItems[j].Name })
	fmt.Println("sorted by name in desending order      :", item.SubItems)


   }