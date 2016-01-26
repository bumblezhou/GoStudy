package main

import (
  "fmt"
  "runtime"
  "time"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("linux")
  default:
    fmt.Printf("%s.\n", os)
  }

  fmt.Printf("When is Saturday?\n")
  today := time.Now().Weekday()
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In 2 days.")
  default:
    fmt.Println("Too far away...")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 18:
    fmt.Println("Good afternoon!")
  default:
    fmt.Println("Good evening!")
  }

  defer fmt.Println("World")
  fmt.Println("Hello")

  fmt.Println("counting....")
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("Done.")
}
