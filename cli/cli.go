package cli

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thekuwayama/details/format"
)

func Run() {
	var summary = flag.String("summary", "", "summary tag")
	var file = flag.String("file", "", "input file path")
	flag.Parse()

	var in *bufio.Scanner
	if file == nil || len(*file) == 0 {
		in = bufio.NewScanner(os.Stdin)
	} else {
		f, err := os.Open(*file)
		if err != nil {
			log.Fatal(fmt.Errorf("file not found: %s", *file))
			os.Exit(1)
		}
		defer f.Close()
		in = bufio.NewScanner(f)
	}

	fmt.Println(format.Header(summary))
	for in.Scan() {
		fmt.Println(in.Text())
	}
	fmt.Println(format.Footer())
}
