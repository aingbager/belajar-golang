package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  create := bufio.NewReader(os.Stdin)
  fmt.Print("masukan kalimat:")
  kalimat,err := create.ReadString('\n')
  if err != nil{
    fmt.Println("error bor",err)
    return
  }
  fmt.Println("kalimat: ",kalimat)
}
