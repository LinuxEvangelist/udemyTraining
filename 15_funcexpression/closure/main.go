package closure

import "fmt"

func Closure(){
	x:=42
	fmt.Println(x)
	{
		fmt.Println(x)
		y:= "The credit belongs with the one who is in the ring"
		fmt.Println(y)
	}
	increment :=wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

func wrapper() func() int {
	//x:=0
	var x int
	return func() int {
		x++
		return x
	}
}