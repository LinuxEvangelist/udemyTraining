package callbacks

import "fmt"

func visit(numbers []int,callback func(int))  {
	for _, n:=range numbers{
		callback(n)
	}
}
//Calling back this function
//func(n int) {
//		fmt.Println(n)
//	})
func Calls()  {
	visit([]int{1,2,3}, func(n int) {
		fmt.Println(n)
	})

	//Another example
	xs := filters([]int{1,2,3,4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
}

func filters(numbers []int,callback func(int) bool) []int {
	xs := []int{}
	for _, n:= range numbers  {
		if callback(n) {
			xs =append(xs,n)
		}
	}
	return xs
}


