package lib

import (
	"bufio"
	"io"
	"fmt"
	"os"
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
