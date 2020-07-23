package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/thekuwayama/details/format"
)

var summary = flag.String("summary", "", "summary tag")

func Run() {
	flag.Parse()

	fmt.Println(format.Header(summary))
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Println(in.Text())
	}
	fmt.Println(format.Footer())
}
