package main

import(
	log "github.com/Sirupsen/logrus"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot find home directory")
	}

	templateFilepath := filepath.Join(workingDirectory, "src/", "template.html")

	template := readFile(templateFilepath)

	fmt.Println(template)
}

func readFile(filepath string) (string){
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Cannot find ", filepath)
	}

	return string(contents)
}