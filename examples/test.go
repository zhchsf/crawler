package main

import(
  "fmt"
  "time"
)

func main(){
  sign := make(chan byte, 1)
  timer := make(chan int, 10)
  go func(){
    for i := 0; i < 10; i++ {
      timer <- i
      time.Sleep(3 * time.Second)
      fmt.Println(i)
    }
    sign <- 1
  }()
  go func(){
    for{
      select {
        case <-timer:
          fmt.Println(111)
      } 
    }
  }()
  <-sign
}