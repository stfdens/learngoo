// package main

// import (
// 	"fmt"
// 	"time"
// )

// type Manusia struct {
// 	Nama string
// 	Age  int
// }

// func (p *Manusia) kirim(ch chan Manusia, orang Manusia) {
// 	p.Nama = orang.Nama
// 	p.Age = orang.Age
// 	ch <- Manusia{p.Nama, p.Age}
// }

// func (p *Manusia) terima(ch chan Manusia) {
// 	data := <-ch
// 	fmt.Println(data)
// }

// func main() {
// 	ch := make(chan Manusia)
// 	orang := Manusia{Nama: "Asep", Age: 16}

// 	go orang.kirim(ch, orang)
// 	go orang.terima(ch)

// 	time.Sleep(time.Second)
// }