package main

import "fmt"

func main()  {
	b := true
	if !true {
		fmt.Println("This did not run")
	}
	if !false{
		fmt.Println("This ran")
	}
	//Inizialitation
	if food := "Chocolate";b{
		fmt.Println(food)
	}

	if false{
		fmt.Println("false")
	}else {
		fmt.Println("Ok")
	}

	if false{
		fmt.Println("false")
	}else if true{
		fmt.Println("ssss")
	}else {
		fmt.Println("Ok")
	}

	for i:=0; i <=100; i++ {
		if i%3 == 0{
			fmt.Println(i)
		}
	}

}
