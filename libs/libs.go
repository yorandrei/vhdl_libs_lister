package libs

import (
	"os"
	"path/filepath"
)

func ListFiles() ([]os.FileInfo, error) {
	//rootDir := "C:\\c\\Frosty\\Repo"

	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(rootDir)
	if err != nil {
		return nil, err
	}

	path, _ := os.Getwd()
	files := getFiles(f, path)

	return files, nil
}

func getFiles(dir *os.File, path string) []os.FileInfo {
	var files []os.FileInfo
	dirs := []*os.File{}

	//fmt.Println("getFiles called from ", dir.Name())

	objects, err := dir.Readdir(0)
	if err != nil {
		return nil
	}

	for _, v := range objects {
		if v.Name()[0] == '.' {
			continue
		}
		if v.IsDir() {
			newpath := filepath.Join(path, v.Name())
			d, err := os.Open(newpath)
			fs := getFiles(d, newpath)
			files = append(files, fs...)
			if err == nil {
				dirs = append(dirs, d)
			}
		} else {
			files = append(files, v)
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
