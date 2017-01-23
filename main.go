package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	gta "github.com/vasili-v/gotauto/lib"
)

func usage() {
	name := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\t%s <template> ...\n", name, name)

	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage
	flag.Parse()
}

func main() {
	templates := flag.Args()
	if len(templates) == 0 {
		fmt.Fprintf(os.Stderr, "no templates specified\n")
		flag.Usage()

		os.Exit(1)
	}

	for _, path := range templates {
		err := gta.InstantiateTemplateByPath(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
}
