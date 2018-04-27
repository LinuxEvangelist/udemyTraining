package main

import "fmt"

//Parameters and arguments, differences
//func receivers(parameters) pass arguments and returns{}
func main()  {
	fmt.Println("Hi")
	greet("ss")
	greet("H")
	greetx("DDD","sss")
	fmt.Println(greeta("Jane","Doe"))
	fmt.Println(greetc("Jane","Doe"))
	n := average(43,56,87,12,45,57)
	fmt.Println(n)
	data := []float64{43,56,87,12,45,57}
	nx := average1(data...)
	fmt.Println(nx)
}

func greet(name string){
	fmt.Println(name)
}

func greetx(fname,lname string)  {
	fmt.Println(fname,lname)
}
func greeta(fname,lname string) string {
	return fmt.Sprint(fname,lname)
}
// return naming
func greetb(fname,lname string) (s string) {
	s = fmt.Sprint(fname,lname)
	return
}
// return multiple values
func greetc(fname,lname string) (string,string) {
	return fmt.Sprint(fname,lname), fmt.Sprint(lname,fname)
}

//variadic
func average(sf ...float64)  float64{
	//cumulating
	total:= 0.0
	for _, v:= range sf  {
		total +=v
	}
	return total / float64(len(sf))
}

//variadic args
func average1(sf ...float64) float64  {
	//func average1(sf []float64) float64  {

		total:= 0.0
	for _,v:= range sf  {
		total += v
	}
	return total /float64(len(sf))
}

func sumLargeNumbers(v string)  {
	x = []byte(v)
	return
}