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

	fmt.Println("call hello funciont")
	str := hello()
	fmt.Println(str)
	fmt.Println("call foobar")
	foobar(16)
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

// foobar divided by 3 foo and divided by 5 bar both foobar else print number
func foobar(n int) {
	fmt.Println("print if divided by 3 foo  divided by 5 bar both foobar \n otherwise print the nubmer")
	for i := 1; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("foobar")
		} else {
			if i%3 == 0 {
				fmt.Println("foo")
			} else if i%5 == 0 {
				fmt.Println("bar")
			} else {
				fmt.Println(i)
			}
		}
	}

}

//hello returns a string simple print
func hello() string {
	return "this is a returned string from a func"

}
