package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//set unqiue seed
	rand.Seed(time.Now().UnixNano())
	fmt.Println("current time is ", time.Now())
	printRandom()
	randomMain()
	fmt.Println("closure in go i.e. nest functions")
	f := testing() //referance the funciton returned
	fmt.Println(f(), f())
	fmt.Println("call hello function")
	str := hello()
	fmt.Println(str)
	fmt.Println("call fizzbuzz")
	fizzbuzz(16)
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomMain() {
	fmt.Println("random between 5 and 50")
	fmt.Println(randomInt(5, 50)) //get an int in the 1...10 range
	fmt.Println("random string: ", randomString(20))
}

func printRandom() {
	fmt.Println("random number printed seed based on time")
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println("--------------------")
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

// fizzbuzz divided by 3 fizz and divided by 5 buzz both foobar else print number
func fizzbuzz(n int) {
	fmt.Println("print if divided by 3 fizz  divided by 5 buzz both foobar \n otherwise print the nubmer")
	for i := 1; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}

}
func testing() func() int {
	a, b := 2, 2
	return func() int {
		x := a + b
		a, b = a+a+a, b+a+b
		fmt.Printf("values of a:%v and b:%v \n", a, b)
		return x
	}
}

//hello returns a string simple print
func hello() string {
	return "Hello there! I hope the day is going great!"
}
