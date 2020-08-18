/*
NOTE: This was rpactice with text template
		NEXT template pracite will be using HTML template!
*/

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	//tpl is template pointer with returned read files data
	tpl, err := template.ParseFiles("index.gohtml")
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
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		//if everything is working and done print success
		fmt.Println("\n success!")
	}

}
