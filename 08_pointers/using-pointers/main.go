package main

import "fmt"

func zero(x int)  {
	x = 0
	fmt.Println(x)
}

/*func main()  {
	x := 5
	zero(x)
	fmt.Println(x) //x is still 5
}
*/
func zero1(z int){
	fmt.Printf("%p\n",&z)
	fmt.Println(&z)
	z = 0
}
/*func main() {
	y := 5
	fmt.Printf("%p\n",&y)
	fmt.Println(&y)
	zero1(y)
	fmt.Println(y)
}*/
//
func zero2(o *int){
	*o = 0
}
func main(){
	h := 5
	zero2(&h)
	fmt.Println(h) // h is 0
}
