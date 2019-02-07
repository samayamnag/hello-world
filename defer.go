package main

import (
  "fmt"
)

func main() {
  test()
}

func mimicError(key string) error {
  return fmt.Errorf("Mimic Error : %s", key)
}

func test() {
  fmt.Println("Start Test")

  err := mimicError("1")

  defer func() {
    fmt.Println("Start Defer")
 
    if err != nil {
      fmt.Println("Defer Error:", err)
    }
  }()

  fmt.Println("End Test")
}