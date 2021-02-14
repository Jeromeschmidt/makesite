package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"
	"strings"
	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"context"
)

func main() {
	saveFile()
}

func readFile(file string) string {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	newfileContents, _ := translateText("en", string(fileContents))

	return string(newfileContents)
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

func translateText(targetLanguage, text string) (string, error) {
        // text := "The Go Gopher is cute"
        ctx := context.Background()

        lang, err := language.Parse(targetLanguage)
        if err != nil {
                return "", fmt.Errorf("language.Parse: %v", err)
        }

        client, err := translate.NewClient(ctx)
        if err != nil {
                return "", err
        }
        defer client.Close()

        resp, err := client.Translate(ctx, []string{text}, lang, nil)
        if err != nil {
                return "", fmt.Errorf("Translate: %v", err)
        }
        if len(resp) == 0 {
                return "", fmt.Errorf("Translate returned empty response to text: %s", text)
        }
        return resp[0].Text, nil
}
