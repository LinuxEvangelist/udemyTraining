package main

import "fmt"

func main()  {
	/*Reference*/
	a:= 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)

	/*Dereference*/
	c:=43
	fmt.Println(c)
	fmt.Println(&c)

	var d *int = &c
	fmt.Println(d)
	fmt.Println(*d)
	fmt.Println(&d)

	/*Using Pointers*/
	x:=43
	fmt.Print("Using Pointers \n")
	fmt.Println(x)
	fmt.Println(&x)

	var y *int =&x
	fmt.Println(y)
	fmt.Println(*y)

	*y = 42
	fmt.Println(y)
	fmt.Println(x)

	/*remainder of division %*/
	f := 13 % 3
	fmt.Println(f)
	if x == 1{fmt.Println("Odd")}else{fmt.Println("Even")}

	for i:=0;i<100;i++{
		if i % 2 == 1{
			fmt.Println("Odd")
		}else{
			fmt.Println("Even")
		}
	}



}
