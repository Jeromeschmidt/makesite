package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"flag"
	"strings"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "File name")
	flag.Parse()
	content := readFile(filename)
	writeFile(filename, content)

}

func readFile(file string) string {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func writeFile(file string, content string) string {
	s := strings.Split(file, ".")

	new_file := "tmp/" + s[0] + ".html"

	err := ioutil.WriteFile(new_file, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	return ""
}
//
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"flag"
// )
//
// func main() {
//
// 	fmt.Println("Hello, world!")
//
// 	file := flag.Parse()
//
// 	temp := readFile(file)
// 	// fmt.Println(temp)
//
// 	writeFile(file)
//
// }
//
// func readFile(file string) string {
// 	fileContents, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return string(fileContents)
// }
//

//
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// )
//
// func main() {
// 	fmt.Println("Hello, world!")
// 	temp := readFile()
// 	fmt.Println(temp)
//
// 	err := ioutil.WriteFile("tmp/first-post.html", []byte(temp), 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// }
//
// func readFile() string {
// 	fileContents, err := ioutil.ReadFile("first-post.txt")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return string(fileContents)
// }
// //
// }

// func renderTemplate()

// package main
//
// import (
//         "html/template"
//         "os"
// 		"io/ioutil"
// 		"log"
//
// )
//
// type entry struct {
//         Name string
//         Done bool
// }
//
// type ToDo struct {
//         User string
//         List []entry
// }
//
// func main() {
//         t := template.Must(template.New("template.tmpl").ParseFiles("new.html"))
//         err = t.Execute(os.Stdout, todos)
//         if err != nil {
//           panic(err)
//         }
//
// 		err := ioutil.WriteFile("tmp/first-post.html", []byte(temp), 0644)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// }
