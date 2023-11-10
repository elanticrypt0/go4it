package go4it

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

func OpenFile(file string) []byte {
	if _, err := os.Stat(file); err != nil {
		log.Fatalln("File dosen't exists: " + file)
	}
	filedata, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Can not open file: " + file)
	}
	return filedata
}

func ReadAndParseToml[T any](file string, stru *T) {
	tomlData := string(OpenFile(file))
	_, err := toml.Decode(tomlData, &stru)
	if err != nil {
		log.Fatalln("Cannot parse file: " + file)
	}
}

func ReadAndParseJson[T any](file string, stru *T) {
	fileData := strings.NewReader(string(OpenFile(file)))
	jsonParser := json.NewDecoder(fileData)
	jsonParser.Decode(&stru)
}

// Gets the current project dir.
func PWD() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// Save or update file
func FileSave(filename string, content []byte) {

}

// Delete file.
func FileDelete(filelocation string) bool {
	return false
}
