package shape

import ("fmt";"math")

type Shape interface {
  area() float64
  perimeter() float64
}

type Rectangle struct {
  length, width float64
}
type Circle struct {
  radius float64
}

func (r Rectangle) area() float64 {
  return (r.length*r.width)
}

func (c Circle) area() float64 {
  return (math.Pi*c.radius*c.radius)
}
func (r Rectangle) perimeter() float64 {
   return (2*(r.length + r.width))
}

func (c Circle) perimeter() float64 {
   return (2*math.Pi*c.radius)
}

func shape()  {

  r := Rectangle{length:5, width:5}
//  fmt.Println("Enter length and width for rectangle: ")
  //fmt.Scanf("%d %d",&r.length,&r.width)
  fmt.Println("Rectangle's length and width: ",r)
  s := Shape(r)
  fmt.Println("Perimeter of rectangle is :",s.perimeter())
  fmt.Println("Area of Rectangle is: ", s.area())

  c := Circle{radius:5}
  fmt.Println("Circle's radius:",c)
  s1 := Shape(c)
  fmt.Println("Perimeter/Circumference of circle is :",s1.perimeter())
  fmt.Println("Area of circle is :", s1.area())

}
