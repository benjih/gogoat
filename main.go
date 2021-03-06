package main

import(
	log "github.com/Sirupsen/logrus"
	"path/filepath"
	"io/ioutil"
	"os"
	"strings"
	"github.com/raymondbutcher/tidyhtml"
	"bytes"
)

var (
	pages = make(map [string]string)
)

func main() {
	log.SetOutput(os.Stdout)


	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot find home directory")
	}

	gogoatSrcDirectory := filepath.Join(workingDirectory, "src/")

	templateFilepath := filepath.Join(workingDirectory, "src/", "template.html")

	filepath.Walk(gogoatSrcDirectory, walk)

	template := readFile(templateFilepath)

	for k, v := range pages {
		filepath := filepath.Join(workingDirectory, k + ".html")
		contents := strings.Replace(template, "$(body)", v, 1)

		var parsedContents bytes.Buffer

		tidyhtml.Copy(&parsedContents, strings.NewReader(contents))

		writeFile(filepath, parsedContents)
	}
}

func readFile(filepath string) (string){
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Cannot find ", filepath)
	}

	return string(contents)
}

func writeFile(filepath string, contents bytes.Buffer) {
	err := ioutil.WriteFile(filepath, contents.Bytes(), 0666)
	if err != nil {
		log.Fatal("Cannot write ", filepath)
	}
}


func walk(path string, f os.FileInfo, err error) error {
	if f.Name() != "template.html" && f.Name() != "src" {
		parts := strings.Split(f.Name(), ".")
		pages[parts[0]] = readFile(path)
	}
	return nil
}