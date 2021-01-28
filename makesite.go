
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello, world!")
	temp := readFile()
	fmt.Println(temp)

	err := ioutil.WriteFile("tmp/first-post.html", []byte(temp), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

// func renderTemplate()
