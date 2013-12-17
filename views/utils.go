package views

import (
	"errors"
	"log"
	"os"
	"path"
	"net/url"
)

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func getTemplate(name string) (string, error) {
	filename := path.Join(templateDir, name)

	if fileExists(filename) {
		return filename, nil
	}
	return "", errors.New("File doesn't exist")
}

func logRequest(method string, from *url.URL, handledBy string) {
	log.Printf("%s  %s  | handler: %s\n", method, from.Path, handledBy)
}
