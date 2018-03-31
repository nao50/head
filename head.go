package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n = flag.Int("n", 10, "lines")
	c = flag.Int("c", 0, "bytes")
)

const BUFSIZE = 1024

func head() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Usage: head [-n lines | -c bytes] [file ...]")
	}

	flag.Parse()
	filename := flag.Args()

	for _, filename := range filename {
		openedfile, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("Can't open file: %s", filename)
		}
		defer openedfile.Close()

		line := bufio.NewScanner(openedfile)
		fmt.Println("==> " + filename + " <==")
		if *c != 0 {
			buf := make([]byte, BUFSIZE)
			for {
				_, err := openedfile.Read(buf)
				if err != nil {
					break
				}
				fmt.Println(string(buf[:*c]))
			}
		} else {
			for j := 1; j < (*n+1) && line.Scan(); j++ {
				fmt.Printf("%d:%s\n", j, line.Text())
			}
		}
	}
	return nil
}

func main() {
	if err := head(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
