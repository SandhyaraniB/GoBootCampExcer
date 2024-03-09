package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func returnOne(x,y string) (string){
return x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
       
	c:=returnOne("Sandhya")
	fmt.Printl(c)

}



