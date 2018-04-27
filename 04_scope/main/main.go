package main

import "fmt"

//package level scope
var x int = 42
//shadow variables, bad recomendation
func max(x int) int{
	return 42+x
}
//Same way we can return a function
func wrapper() func() int{
	x :=0
	return func() int {
		x++
		return x
	}

}

func main()  {
	fmt.Println(x)
	foo()
	//block level scope
	y:= 17
	fmt.Println(y)
	//shadow variables
	max := max(4)
	fmt.Println(max)
	//examples of scopes
	c := 42
	fmt.Println(c)
	{
		fmt.Println(c)
		y := "example de outer and inner scopes {}"
		fmt.Println(y)
	}

	u :=0
	//Anonymous function
	increment := func() int{
		u++
		return u
	}
	fmt.Println(increment())
	fmt.Println(increment())
	increment1 := wrapper()
	fmt.Println(increment1())
	fmt.Println(increment1())

	//blank identifiers
	_ :="helloooo"
}

func foo(){
	fmt.Println(x)
	//fmt.Println(y)
}
