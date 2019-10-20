package main

import "fmt"

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string){
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang")
//	d.p.Speak()
}

func (d *Dog) SpeakTo(host string){
	d.Speak()
	fmt.Println(" ", host)
//	d.p.SpeakTo(host)
}

func main(){
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
