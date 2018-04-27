package main

import "fmt"

//declaration
//var b string
//assingment
//b:= "cowboy"

func main() {
	/*shorthand: Declare,assign,initialize*/
	a := 10
	b := "golang"
	c := 4.17
	d := true
	//Declaring something to zero
	var e int
	var f string
	var g float64
	var h bool
	fmt.Printf("%v \n", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%T \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%T \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%T \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)


}
