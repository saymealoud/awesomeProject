package main

//func main() {
//	fibonacci := Fibonacci(10)
//	fmt.Println(fibonacci)
//
//	//TestFibonacci()
//}

func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {

		return 0
	}
	if n == 1 {

		return 1

	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
