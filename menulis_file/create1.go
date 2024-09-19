package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file_txt, err := os.Create("asu.txt")
	if err != nil {
		fmt.Println("error brow", err)
		return
	}
	defer file_txt.Close() //sesudah dibuat di close

	writters := bufio.NewWriter(file_txt)

	writters.WriteString("ini baris ketiga \n")
	writters.WriteString("ini baris kedua \n")

	writters.Flush()

	fmt.Println("file berhasil dibuat")

	//untuk membaca file yang baru dibuat otomatis tertutup
	//oleh sebab itu harus di buka(open)

	file_txt, err = os.Open("asu.txt")
	if err != nil {
		fmt.Println("file gagal di buka", err)
		return
	}
	defer file_txt.Close()

	data, err := io.ReadAll(file_txt)
	if err != nil {
		fmt.Println("gagal membaca file", err)
		return
	}

	fmt.Println("isi file: ", string(data))

}
