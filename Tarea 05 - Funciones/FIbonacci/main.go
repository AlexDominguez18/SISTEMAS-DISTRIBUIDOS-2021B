package main

import "fmt"

func fibonacci(n int64) int64{
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main (){
	var n int64
	
	fmt.Scan(&n)
	s := make([]int64, n)

	fmt.Println("n = ", n)

	for i := 0; int64(i) < n ; i++ {
		s[i] = fibonacci(int64(i))
	}

	fmt.Println(s)
}