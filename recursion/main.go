package main

import "fmt"

func main() {
	//call func fib, hunio, permutate
	fmt.Println("Hello! \n First function is factorial")
	factorial(4, 1)
	fmt.Println("next the Fibonacci sequence print the 15th in the sequence ")
	fib(15, 0, 1)
	fmt.Println("Next to run is Tower of Hanoi! The number of discs in the puzzle is: 4")
	hannoi(4, "A", "B", "C")
	fmt.Println("last but not least Alexander Bogomolny’s UnOrdered Permutation Algorithm \n For this we will get all permutations from 1-6")
	n := 6
	arr := []int{}
	abp(arr, n, 0, -1)
	fmt.Println("all done!")
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
func fib(n, a, b int) {
	if n == 0 {
		fmt.Println(a)
		fmt.Println("The number is: ", b)
	} else if n == 1 {
		fmt.Println(b)
		fmt.Println("The number is: ", b)
	} else {
		fmt.Printf("%v, ", a)
		n := n - 1
		next := a + b
		cur := b
		fib(n, cur, next)
	}
}

//Tower of Hanoi
//print out each move for n numbers
// f from  , h helper , t target
// move is a pure function that just prints the move to be made
func move(f, t string) {
	fmt.Printf("move disc from %q to %q \n", f, t)
}

func hannoi(n int, f, h, t string) {
	if n == 0 {
		//do nothing there may be a better way to wrtie this part but it works this way sooooooo idk
	} else {
		hannoi(n-1, f, t, h)
		move(f, t)
		hannoi(n-1, h, f, t)
	}
}

//Alexander Bogomolny’s UnOrdered Permutation Algorithm
// show all possible from 1 to n ex: input 2 output 1 2   2 1
func print(perm []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v, ", perm[i])
	}
	fmt.Printf("\n")
}

// n number l level

func abp(perm []int, n, k, lvl int) {
	if len(perm) == 0 {
		for i := 1; i <= n; i++ {
			perm = append(perm, 0)
		}
	}
	// Assign level to zero at start.
	lvl = lvl + 1
	perm[k] = lvl
	if lvl == n {
		print(perm, n)
	} else {
		for i := 0; i < n; i++ {
			if perm[i] == 0 {
				abp(perm, n, i, lvl)
			}
		}
	}
	lvl = lvl - 1

	perm[k] = 0

}
