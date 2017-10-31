/*** author 
     BGH25333
 ***/
     
package main

import(
"fmt"
"math"
)


type shape interface {
    area() float64
    perimeter() float64
}

type rectangle struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rectangle) area() float64 {
    return r.width * r.height
}
func (r rectangle) perimeter() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}


func main() {

    var width float64
    var height float64
    var radius float64
    fmt.Println("Enter the width, height of rectangle and radius of circle : ")

    fmt.Scan(&width,&height,&radius)
 
    r := rectangle{width: width, height: height }
    c := circle{radius: radius}

  
   fmt.Println(" Area of rectangle is :" ,r.area(), "\n Perimeter of rectangle :", r.perimeter())
   fmt.Println(" Area of circle is :" ,c.area(), "\n Perimeter of circle :", c.perimeter())

}