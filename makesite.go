
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, world!")
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

// func renderTemplate()
