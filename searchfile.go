/*** author 
     BGH25333
 ***/
     

package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() () {
    
    var fileName string
    fmt.Print("Enter the file name you want to search : ")
    fmt.Scan(&fileName)

 SearchFile(fileName)
}

func SearchFile(input string){

    searchDir, _ := os.Getwd()
    fileList := []string{}

    fmt.Println("current directory is : ",searchDir)
    fmt.Println("")

   filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileList = append(fileList, path)
        return nil
    })

    for _, file := range fileList {

    	if(strings.Contains(file,input)){
        fmt.Println("Absolute path of file is :", file)
    }
 
    }

}