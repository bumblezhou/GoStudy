package main

import(
  "fmt"
  "math"
)

func sqrt(x float64) string{
  if(x < 0){
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func pow(x, y, lim float64) float64 {
  if v := math.Pow(x, y); v < lim {
    return v
  }
  return lim
}

func my_sqrt(x float64) float64 {
  z := float64(1)
  for i := 0; i < 10; i++ {
    z = z - ((z * z - x) / (2 * z))
    fmt.Printf("iterate %v with the result: %v\n", i, z)
  }
  return z
}

func main() {
  sum := 0
  for i := 0; i < 100; i++ {
    sum += i
  }
  fmt.Printf("%v \n", sum)

  n := 1
  for (n < 1000) {
    n += n
  }
  fmt.Printf("%v \n", n)

  fmt.Println(sqrt(2), sqrt(-4))

  fmt.Println(pow(2, 3, 10))
  fmt.Println(pow(3, 3, 20))
  fmt.Println(pow(2, 5, 100))

  fmt.Printf("my_sqrt(2) is %v\n", my_sqrt(2))
  fmt.Printf("math.Sqrt(2) is %v\n", math.Sqrt(2))
}
