package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n = flag.Int("n", 10, "lines")
)

func head() error {
	if len(os.Args) < 1 {
		return fmt.Errorf("Usage: head [-n lines | -c bytes] [file ...]")
	}

	flag.Parse()
	filename := flag.Args()

	openedfile, err := os.Open(filename[0])
	if err != nil {
		return fmt.Errorf("Can't open file: %s", filename)
	}
	defer openedfile.Close()

	line := bufio.NewScanner(openedfile)
	for i := 1; i < (*n+1) && line.Scan(); i++ {
		fmt.Printf("%d:%s\n", i, line.Text())
	}

	return line.Err()
}

func main() {
	if err := head(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

/*
NaonoiMac:head nao$ head --help
head: illegal option -- -
usage: head [-n lines | -c bytes] [file ...]
*/
