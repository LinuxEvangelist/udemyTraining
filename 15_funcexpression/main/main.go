package main

import "fmt"
import (
	"github.com/LinuxEvangelist/udemyTraining/15_funcexpression/closure"
	"github.com/LinuxEvangelist/udemyTraining/15_funcexpression/callbacks"
	"github.com/LinuxEvangelist/udemyTraining/15_funcexpression/recursion"
	"github.com/LinuxEvangelist/udemyTraining/15_funcexpression/deferxa"
	"github.com/LinuxEvangelist/udemyTraining/15_funcexpression/passbyvalue"
	"github.com/LinuxEvangelist/udemyTraining/15_funcexpression/anonymous"
	"github.com/LinuxEvangelist/udemyTraining/15_funcexpression/boolian"
)

/*func greeting()  {
	fmt.Println("Hi")
}
*/
//func_expression, only way express one function inside function
func main()  {
	greeting := func() {
		fmt.Println("hi")
	}
	greeting()
	fmt.Printf("%T\n",greeting)

	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n",greet)
	//Execute closure
	closure.Closure()

	//Execute calls
	callbacks.Calls()
	fmt.Println("Recursion--> ",recursion.Factorial(3))

	//Execute defer
	deferxa.Deferr()

	//Execute passbyvalue
	fmt.Println("***Examples passbyvalue***")
	passbyvalue.Passby()
	passbyvalue.PassbyValueMap()

	//Execute anonymous self-execution function
	anonymous.AnonymousSelfExecuting()

	//Execute boolean expression
	fmt.Println("***Examples Boolean Func***")
	boolian.BooleanT()
}
//-->return and function
func makeGreeter() func() string {
	return func() string {
		return "Another hi"
	}

}