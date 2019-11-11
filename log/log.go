package main
import (
	_ "fmt"
	"log"
)

func main(){
	log.Printf("%d", "logger...\n")
	log.Fatal("logger fatal...")
}
