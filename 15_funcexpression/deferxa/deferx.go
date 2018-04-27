package deferxa

import "fmt"
//use defer for execute this last
func Deferr()  {
	defer world()
	hello()
}

func hello(){
	fmt.Print("hello ")
}
func world()  {
	fmt.Println("world")
}
