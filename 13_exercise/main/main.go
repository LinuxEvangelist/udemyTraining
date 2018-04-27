package main

import "fmt"

const p  = "death & no"


//go fmt ./...(format files)
func main() {
	name := "Ricardo"
	fmt.Println("Hello World", name)

	var nameenter string
	fmt.Print("Enter Name: ")
	fmt.Scan(&nameenter)
	fmt.Println("Hello ", nameenter)

	var sn, ln int32
	fmt.Print("Enter Small Number: ")
	fmt.Scan(&sn)
	fmt.Print("Enter Larger Number: ")
	fmt.Scan(&ln)
	fmt.Println("Remainder: ", ln/sn)

	for x := 0; x <= 100; x++ {
		//fmt.Println(x)
		if (x%2 != 0) {
			continue
		}
		fmt.Println("Even Number", x)
	}

	for y:= 1; y <=100;  y++{
		if y%15 == 0{
			fmt.Println(y," FizzBuzz")
		}else if y%5 == 0{
			fmt.Println(y," Buzz")
		//}else if y%3 == 0 && y%5 == 0{
		}else if y%3 == 0{
			fmt.Println(y," Fizz")
		}else{
		fmt.Println(y)
		}
	}
	sum :=0
	for t:=0; t < 1000; t++ {
		if t%3==0 || t%5==0{
			fmt.Println( "Test Numbers: ",t)
			sum += t
		}
	}
	fmt.Println("Sum all-->:",sum)

}
