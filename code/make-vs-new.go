package main

import "fmt"

func main(){

var i *int
i = new(int)
fmt.Printf("> %d\n", *i)

var s []int = make([]int, 10)

fmt.Println("> ", s)


}