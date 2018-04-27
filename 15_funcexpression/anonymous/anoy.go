package anonymous

import "fmt"

func AnonymousSelfExecuting()  {
	func(){
		fmt.Println("Executing this")
	}()
}
