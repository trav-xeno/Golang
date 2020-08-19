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
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&name)
	fmt.Println("Please enter your favorite food: ")
	foodReader := bufio.NewReader(os.Stdin) //create new reader
	food, readErr := foodReader.ReadString('\n')
	if readErr != nil {
		log.Fatalln(readErr)
		fmt.Println("an error occured while read input!")
	}
	fmt.Println("Please enter a fun fact: ")
	funFactReader := bufio.NewReader(os.Stdin) //create new reader
	fact, factErr := funFactReader.ReadString('\n')
	if factErr != nil {
		log.Fatalln(factErr)
		fmt.Println("an error occured while read input!")
	}

	if err != nil {
		log.Fatalln(err)
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
	user := person{Name: name, FavFood: food, FunFact: fact}
	fmt.Println("name: " + user.Name + "\n favorite food: " + user.FavFood + "\n fun fact: " + user.FunFact)
	err = tpl.Execute(os.Stdout, user)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create("./index.html")
	if err != nil {
		log.Println("create file: ", err)
	}
	err = tpl.Execute(f, user)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()

}
