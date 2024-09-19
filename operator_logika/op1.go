package main

import "fmt"

func main() {
  //operator logika
  var left = false
  var right = true

  var letfAndRight = left && right
  fmt.Println(letfAndRight)

  var leftOrright = left || right
  fmt.Println(leftOrright)

  var leftreverse = !left
  fmt.Println(leftreverse)
}
