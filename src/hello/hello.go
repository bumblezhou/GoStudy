package main

import (
  "fmt"
  "stringutil"
  "time"
  "math/rand"
  "math"
  "math/cmplx"
)

func add(x, y int) int {
  return x + y
}

func swap(x, y string) (string, string){
  return y, x
}

func split(sum int) (x, y int){
  x = sum * 4 / 9
  y = sum - x
  return
}

var (
  ToBe bool = false
  MaxInt uint64 = 1<<64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

var n, m int = 1, 2

func main() {
  fmt.Printf(stringutil.Reverse("\nHello, world!\n"))
  fmt.Println("The time is:", time.Now())
  fmt.Println("My favorite number is:", rand.Intn(100))
  fmt.Printf("Now you get %g problems. and the value of PI is: %g.\n", math.Sqrt(7), math.Pi)
  fmt.Println("23 + 45 = ", add(23, 45))

  //a, b := swap("world", "hello")
  //fmt.Println(a, b)

  fmt.Println(split(17))

  var c, python, java = false, false, "No!"
  k := 3
  fmt.Println(c, java, python, n, m, k)

  //const f = "%T(%v)\n"
  //fmt.Printf(f, ToBe, ToBe)
  //fmt.Printf(f, MaxInt, MaxInt)
  //fmt.Printf(f, z, z)

  //var i int
  //var f float64
  //var b bool
  //var s string
  //fmt.Printf("%v %v %v %q\n", i, f, b, s)

  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Printf("%v %v %v %v\n", x, y, f, z)

  v := 42
  fmt.Printf("v is type of %T\n", v)
}
