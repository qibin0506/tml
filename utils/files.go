package utils

import (
	"os"

	"github.com/PuerkitoBio/goquery"
)

func ReadFile(filepath string) *os.File {
	contents, err := os.Open(filepath)
	CheckError(err)

	return contents
}

func ReadFileAsDom(filepath string) *goquery.Document {
	contents := ReadFile(filepath)

	dom, error := goquery.NewDocumentFromReader(contents)
	CheckError(error)

	return dom
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func CreateDir(name string) {
	if FileExists(name) {
		return
	}

	for i := 0; i < 3; i++ {
		if err := os.MkdirAll(name, 0777); err == nil {
			break
		}
	}
}