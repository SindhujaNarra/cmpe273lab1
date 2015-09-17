package main

import "fmt"

func main() {
 var inp int64 = 0
 var x int64 = 0
 fmt.Println("Enter a number")
 fmt.Scanf("%d",&inp)
 x= fib(inp)
 fmt.Println("Fib is: ",x)
}

func fib(n int64) int64 {
  //first := 0
  //second := 1
  if n == 0 {
    return 0
  }else if (n == 1 || n == 2) {
     return 1
  }else {
    return (fib(n-1)+fib(n-2))
  }
}
