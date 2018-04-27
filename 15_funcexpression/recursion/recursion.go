package recursion

func Factorial(x int)int{
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}


//3*2*1