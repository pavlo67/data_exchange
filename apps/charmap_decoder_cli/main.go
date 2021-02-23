package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {

	f, e := os.Open("example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	decoder := charmap.Windows1251.NewDecoder()
	reader := decoder.Reader(f)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//// Запись строки в кодировке Windows-1252
	//encoder := charmap.Windows1252.NewEncoder()
	//s, e := encoder.String("This is sample text with runes Š")
	//if e != nil {
	//	panic(e)
	//}
	//ioutil.WriteFile("example.txt", []byte(s), os.ModePerm)

}
