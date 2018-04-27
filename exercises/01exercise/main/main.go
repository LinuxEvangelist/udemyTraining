package main

import "fmt"
import "math/big"

func main()  {
	r,x := returnint(2)
	fmt.Println(r,x)

	//func expression, only way declaration function inside function
	rx := func(n int) (v int,v2 bool){
		return n/2,n%2 == 0
	}
	a,b := rx(2)
	fmt.Println(a,b)

	//Variadic
	data := []int{43,56,87,100,45,57}
	nx := variadic(data...)
	fmt.Println(nx)

	//Value of expression
	fmt.Printf("%v\n", (true && false)||(false&&true)||!(false&&false))
	foo()
	foo(1,2,3)
	aSlice := []int{1,2,3}
	foo(aSlice...)

	//Euler sum bignumbers
	//var gd uint64
	var gd big.Int
	//gd = 37107287533902102798797998220837590246510135740250
	//gd.SetInt64(37107287533902102798797998220837590246510135740250)
	fmt.Printf("%v\n",&gd)
	fmt.Printf("%T\n",gd)
}


func returnint(x int)(r1 int,rs2 bool) {
	r1 = x/2;
	//return r1
	if x%2 == 0{
		rs2 = true
	}else{
		rs2 = false
	}
	/*
	return n / 2, n%2 == 0
	*/
	return r1,rs2
}

func variadic(nx ...int) (m int){
	//r1 := 0

	//if len(nx) > 0 {
		m = nx[0]
	//}
	/*for i:=0; i < len(nx); i++ {
		if nx[i] > m  {
			m = nx[i]
		}
	}*/
	//greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
	for  _, v:= range nx  {
		if v > m {
			m = v
		}
	}

	return m
}

func foo(vl ...int)  {
	fmt.Println(vl)
}