package libs

import (
	"os"
	"path/filepath"
)

func ListFiles() ([]string, error) {
	//rootDir := "C:\\c\\Frosty\\Repo"

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

	//fmt.Println("getFiles called from ", dir.Name())

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

/*
func GetLibraries(files []os.FileInfo) []string {
	libraries := []string{}
	for _, file := range files {
		file.
	}
	return libraries
}
*/
