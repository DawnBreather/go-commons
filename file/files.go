package file

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func FindFiles(path, pattern string) []os.FileInfo {
	libRegEx, e := regexp.Compile(pattern)
	if e != nil {
		log.Fatal(e)
	}

	var files []os.FileInfo

	e = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			files = append(files, info)
			//println(info.Name())
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}

	return files
}
