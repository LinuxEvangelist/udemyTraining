package boolian

import "fmt"

//expressions(horizontal), statements()
func BooleanT()  {
	truefalse()

	//same way
	if true || false{
		fmt.Println("This ran with conditional OR ||")
	}
	if true && false{
		fmt.Println("This ran with conditional AND &&")
	}
}


func truefalse()  {
	if true{
		fmt.Println("This ran")
	}
	if false {
		fmt.Println("This did not run")
	}
}
