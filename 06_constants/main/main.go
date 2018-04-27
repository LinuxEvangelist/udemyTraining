package main

import "fmt"

const p  = "death & no"

const (
	Pi = 3.14
	Language ="Go"
)
//Iota
const (
	/*A = iota //0
	B = iota //1
	C = iota //2*/

	A = iota //0
	B   //1
	C   //2
)
//Iota reset count
const(
	D = iota //0
	E   //1
	F   //2
)

const(
	_ = iota //0
	G = iota *10  //1*10
	H = iota *10  //2*10
)

const(
	_ = iota
	//swifting bits
	KB = 1 << (iota *10) // 1<< (1*10)
	MB = 1 << (iota *10)// 2<< (1*10)
	GB = 1 << (iota *10)// 3<< (1*10)
	TB = 1 << (iota *10)// 4<< (1*10)
)

func main()  {
	const q  = 42
	fmt.Println("p - ",q)
	fmt.Println("p - ",p)
	fmt.Println(Pi)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(G)
	fmt.Println(H)

	fmt.Printf("%d\t",KB)
	fmt.Printf("%b\n",KB)
	fmt.Printf("%d\t",MB)
	fmt.Printf("%b\n",MB)
	fmt.Printf("%d\t",GB)
	fmt.Printf("%b\n",GB)
	fmt.Printf("%d\t",TB)
	fmt.Printf("%b\n",TB)

	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
}
