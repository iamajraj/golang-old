package main

import "fmt"

func init(){
	fmt.Println("THIS WILL RUN WITHOUT CALLING")
}

func main() {
	fmt.Println("HELLO")
}