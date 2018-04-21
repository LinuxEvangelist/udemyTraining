package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(42)
	//On Binary - variadic
	fmt.Printf("%d - %b \n",42,42)
	//hexadecimal
	fmt.Printf("%d - %b - %x \n" ,42,42,42)
	//fmt.Printf("%d - %b - %#x \n" ,42,42,42)
	fmt.Printf("%d \t %b \t %#x \n" ,42,42,42)
	/*Loop*/
	for i:=0;i<200;i++{
		fmt.Printf("%d \t %b \t %x \n",i,i,i)
	}
	/*UTF-8*/
	for i:=0;i<200 ;i++  {
		fmt.Printf("%d \t %b \t %#x \t %q \t %s \n",i,i,i,i,i)
	}
}