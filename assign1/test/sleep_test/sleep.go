package sleep

import (
  "time"
)

func sleep() {
   <-time.After(time.Second*3)
}

/*func main() {

  fmt.Println("Hello")
  sleep()
  fmt.Println("Hi")
  sleep()
  fmt.Println("How are you?")
  sleep()
  fmt.Println("Have a good day")
}*/
