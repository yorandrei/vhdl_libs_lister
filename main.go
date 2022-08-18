package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yorandrei/vhdl_libs_lister/libs"
)

func main() {
	start := time.Now()
	files, err := libs.ListFiles()
	t := time.Now()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("ListFiles took: ", t.Sub(start))

	/*
		fmt.Println("Inspected Source Files:")
		for _, f := range files {
			fmt.Println(f)
		}
		fmt.Printf("\n\n\n")
		fmt.Println("Found Libraries:")
	*/

	start = time.Now()
	libraries := libs.GetLibraries(files)
	t = time.Now()
	for _, l := range libraries {
		fmt.Println(l)
	}
	fmt.Println("Searching libraries took: ", t.Sub(start))

}
