package main

import (
  "fmt"
  "math/rand"
  "time"
)


func main() {
  n := 30
  arr := generate(n)

  fmt.Println(arr)
  shuffle(arr)
  fmt.Println(arr)
}

func generate(n int) []string {
  arr := make([]string, n)
  for i := range arr {
    arr[i] = fmt.Sprintf("%02d", i)
  }
  return arr
}

func shuffle(data []string) {
  rand.Seed(time.Now().UnixNano())
  n := len(data)
  for i := n -1; i >= 0; i-- {
    j := rand.Intn(i + 1)
    data[i], data[j] = data[j], data[i]
  }
}
