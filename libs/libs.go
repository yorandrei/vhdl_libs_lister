package libs

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func ListFiles() ([]string, error) {

	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(rootDir)
	if err != nil {
		return nil, err
	}

	files := getFiles(f, rootDir)

	return files, nil
}

func getFiles(dir *os.File, path string) []string {
	var files []string

	objects, err := dir.Readdir(0)
	if err != nil {
		return nil
	}

	for _, v := range objects {
		if v.Name()[0] == '.' {
			continue
		}
		fullpath := filepath.Join(path, v.Name())
		if v.IsDir() {
			d, err := os.Open(fullpath)
			defer d.Close()
			if err == nil {
				fs := getFiles(d, fullpath)
				files = append(files, fs...)
			}
		} else {
			files = append(files, fullpath)
		}
	}

	return files
}

func GetLibraries(files []string) []string {
	libraries := []string{}
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s := scanner.Text()
			if strings.Contains(s, "use ") {
				spl := strings.Split(s, " ")
				if spl[0] == "use" && len(spl) == 2 {
					lib := spl[1]
					if !contains(libraries, lib) {
						libraries = append(libraries, lib)
					}
				}
			}
		}
	}
	return libraries
}

func FindLibs(files <-chan string, libs chan<- string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	for filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s := scanner.Text()
			if strings.Contains(s, "use ") {
				spl := strings.Split(s, " ")
				if spl[0] == "use" && len(spl) == 2 {
					lib := spl[1]
					libs <- lib
				}
			}
		}
	}
}

func FilterLibs(libs <-chan string, flibs *([]string)) {
	for l := range libs {
		if !contains(*flibs, l) {
			*flibs = append(*flibs, l)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
