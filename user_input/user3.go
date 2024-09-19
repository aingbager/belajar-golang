package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	baca := bufio.NewReader(os.Stdin) // menulis kalimat panjang
	fmt.Print("masukan kalimat: ")
	kalimat, err := baca.ReadString('\n') //baca kalimat panjang
	if err != nil{
	  fmt.Println("terjadi kesalahan...",err)
	  return
	}
	fmt.Println("kalimat yang anda masukan:", kalimat)
}
