package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"flag"
	"strings"
)

func main() {
	// var filename string
	// var dir string
	// flag.StringVar(&filename, "file", "", "File name")
	// flag.StringVar(&dir, "dir", "", "Directory name")
	// flag.Parse()
	//
	// content := readFile(filename)
	// writeFile(filename, content)
	saveFile()
}

func readFile(file string) string {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func writeFile(file string, content string) {
	s := strings.Split(file, ".")

	new_file := "tmp/" + s[0] + ".html"

	err := ioutil.WriteFile(new_file, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func saveFile() {
	var filename string
	var dir string
	flag.StringVar(&filename, "file", "", "File name")
	flag.StringVar(&dir, "dir", "", "Directory name")
	flag.Parse()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
	    log.Fatal(err)
	}
	for _, f := range files {
		if f.Name()[len(f.Name())-4:] == ".txt"{
			content := readFile(dir + "/" + f.Name())
			writeFile(f.Name(), content)
		}
	}
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
