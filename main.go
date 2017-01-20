package main

import (
	"bufio"
	"io"
	"flag"
	"fmt"
	"os"
	"path"
)

func InstantiateTemplate(r io.Reader) error {
	sum := 0
	number := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		number += 1
		line := s.Text()
		sum += len(line)
	}

	fmt.Printf("Average: %0.2f spl\n", float64(sum)/float64(number))

	return nil
}

func InstantiateTemplateByPath(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return InstantiateTemplate(f)
}

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
		err := InstantiateTemplateByPath(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
}
