package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	//Nested Loops
	for j := 0; j < 10; j++ {
		for k := 0; k < 10; k++ {
			fmt.Println(j, " - ", k)
		}

	}
	//With some conditions, like while command
	ii := 0
	for ii < 10 {
		fmt.Println(ii)
		ii++
	}
	//without conditions
	/*iii := 0
	for {
		fmt.Println(iii)
		iii++
	}*/
	//Another way,break
	u := 0
	for {
		fmt.Println(u)
		if u >= 10 {
			break
		}
		u++
	}
	//continue
	d := 0
	for {
		d++
		if d%2 == 0 {
			//stop iteration and return to the loop, for another value
			continue
		}
		fmt.Println(d)
		if d >= 50 {
			break
		}
	}
	//loops over, rune(is a single character '' alias for int32(or 4 bytes) not strings("" ``),
	// literal values
	fmt.Println([]byte("Hello Go"))
	for i:=50; i <=140;i++  {
		fmt.Println(i," - ", string(i), " - ", []byte(string(i)))
		fmt.Printf("%v - %v - %v \n",i, string(i), []byte(string(i)))

	}
	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n",foo)
	fmt.Printf("%T \n", rune(foo))

}
