package main

import "fmt"

const metersToYards float64  = 1.09361
func main()  {
	a:=43
	fmt.Println("a - ",a)
	fmt.Println("a'ss memory address - ",&a)
	fmt.Printf("%d\t",&a)
	//
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards :=meters*metersToYards
	fmt.Println(meters," meter is ",yards, " yards.")


}
