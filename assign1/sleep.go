package main

import ("fmt"
  "time"
)

func sleep() {
   <-time.After(time.Second*3)  // using left arrow since time.After returns a channel
  }

func main() {

  fmt.Println("Hello")
  sleep()
  fmt.Println("Hi")
  sleep()
  fmt.Println("How are you?")
  sleep()
  fmt.Println("Have a good day")
}
