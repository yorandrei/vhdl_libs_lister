package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/yorandrei/vhdl_libs_lister/libs"
)

func main() {
	const NUM_THREADS = 10
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
	fch := make(chan string)
	lch := make(chan string)

	var libraries []string

	var mu sync.Mutex
	var wg sync.WaitGroup
	start = time.Now()

	for n := 0; n < 3; n++ {
		go libs.FilterLibs(lch, &libraries, &mu)
	}

	for t := 0; t < NUM_THREADS; t++ {
		wg.Add(1)
		go libs.FindLibs(fch, lch, &wg)
	}

	for _, f := range files {
		fch <- f
	}
	close(fch)
	wg.Wait()
	close(lch)
	t = time.Now()

	fmt.Println("Searching libraries took: ", t.Sub(start))

	start = time.Now()
	for _, l := range libraries {
		fmt.Println(l)
	}
	t = time.Now()

	fmt.Println("Printing took: ", t.Sub(start))
}
