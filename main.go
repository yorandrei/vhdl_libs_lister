package main

import (
	"fmt"
	"log"
	"vhdl/libs"
)

func main() {
	files, err := libs.ListFiles()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Inspected Source Files:")
	for _, f := range files {
		fmt.Println(f)
	}
	/*
		fmt.Printf("\n\n\n")
		fmt.Println("Found Libraries:")
		libraries := libs.GetLibraries(files)
		for _, l := range libraries {
			fmt.Println(l)
		}
	*/
}
