/*
NOTE: This was rpactice with text template
		NEXT template pracite will be using HTML template!
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/template"
)

type person struct {
	Name    string
	FavFood string
	FunFact string
}

func main() {
	//tpl is template pointer with returned read files data
	tpl, err := template.ParseFiles("index.gohtml")
	var name, food, fact string
	//print questions to user
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&name) //single word input only but I hsould add a check to make sure they only enter one word
	fmt.Println("Please enter your favorite food: ")
	// take multiple words input
	foodReader := bufio.NewReader(os.Stdin) //create new reader
	food, readErr := foodReader.ReadString('\n')
	if readErr != nil {
		log.Println("error getting food data: ", readErr)
	}
	fmt.Println("Please enter a fun fact: ")
	funFactReader := bufio.NewReader(os.Stdin) //create new reader
	fact, factErr := funFactReader.ReadString('\n')
	if factErr != nil {
		log.Println("error getting fun fact data: ", factErr)
	}

	//this part reads template and prints to commandline and takes in data to add to webpage
	// or
	//if multiple templates use ExecuteTemplate
	// create a file
	/* nf , createFilleErr := os.create("index.html")
		if createFileErr != nil{
			log.Fatalln(createFileErr)
		}
		defer nf.close()
		//still needs template becuase it basially is a pointer to collection oif templates
	    err = tpl.Execure(nf, nil) //take in no data
	*/
	//create user object based on input
	user := person{Name: name, FavFood: food, FunFact: fact}
	aiOne := person{Name: "aiOne", FavFood: "good bytes ;)", FunFact: "I wont take over the world"}
	mathis := person{Name: "Mathis the dog", FavFood: "Turkey", FunFact: "I guide the blind."}
	blob := person{Name: "The Blob", FavFood: "raspberry", FunFact: "I play trumpet and swing dance!"}
	//make a list of people to add to the carousel on the html page
	people := []person{user, aiOne, mathis, blob}
	fmt.Println(people)
	//fmt.Println("name: " + user.Name + "\n favorite food: " + user.FavFood + "\n fun fact: " + user.FunFact)
	err = tpl.Execute(os.Stdout, user)
	if err != nil {
		log.Println("error occured", err)
	}
	f, err := os.Create("./index.html")
	if err != nil {
		log.Println("create file error: ", err)
	}
	err = tpl.Execute(f, user)
	if err != nil {
		log.Println("execute error: ", err)
		return
	}
	f.Close()

}
