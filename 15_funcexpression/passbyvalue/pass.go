package passbyvalue

import "fmt"

func Passby()  {
	age := 44
	fmt.Println(&age) //1.0xc4200141f8(memory address)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *int)  { //pointers
	fmt.Println(z) //2. 0xc4200141f8
	fmt.Printf("%T\n",z) //*int
	fmt.Println(*z) //44  //derefence
	*z = 24
	fmt.Println(z) //0xc4200141f8
	fmt.Println(*z) //24
}

//String are inmutable,primitive value--> not a reference type
func PassbyValue()  {
	//slides []string, is a reference type
	m1 := make([]string,1,25)
	fmt.Println(m1) //[ ]
	changeMe1(m1)
	fmt.Println(m1) //[Todd]
}

func changeMe1(z []string)  {
	z[0] = "Todd"
	fmt.Println(z) //[Todd]
}
func PassbyValueMap()  {
	x:= make(map[string]int)
	changeMe2(x)
	fmt.Println(x["Tood"]) //44
}
func changeMe2(z map[string]int)  {
	z["Todd"] = 44
}



