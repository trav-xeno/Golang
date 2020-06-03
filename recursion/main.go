package main

import "fmt"

func main() {
	//call func fib, hunio, permutate
	fmt.Println("Hello! \n First function is factorial")
	factorial(4, 1)
	fmt.Println(" nect the Fibonacci sequence")
	fib(6)
	fmt.Println("Next to run is Tower of Hanoi ")
	hannoi(5)
	fmt.Println("last but not least Alexander Bogomolny’s UnOrdered Permutation Algorithm ")
	permutate(5)
	fmt.Println("all done for now next your input!")
	//call input fucntions

}

//factorial
// n(n - 1)
func factorial(in, fact int) {
	if in == 1 {
		fmt.Println("factorial result: ", fact)
	} else {
		next := in - 1
		fact := in * fact
		factorial(next, fact)
	}
}

//Fibonacci recursion function
// take the number you wish to see  example 50 will show up to 50th number in serise
func fib(in int) {
	fmt.Println(in)

}

//Tower of Hanoi
//print out each move for n numbers
func hannoi(num int) {
	fmt.Println(num)
}

//Alexander Bogomolny’s UnOrdered Permutation Algorithm
// show all possible from 1 to n ex: input 2 output 1 2 2 1
func permutate(num int) {
	fmt.Println(num)
}
