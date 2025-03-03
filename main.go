package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bodagovsky/lightspeed-test/counter"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("error encountered while executing program: %s\n", err.(error).Error())
			os.Exit(1)
		}
	}()

	if len(os.Args) < 2 {
		panic(errors.New("no filename provided for ipV4 addresses to count"))
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ctr := counter.NewIpAdressCounter()
	r := bufio.NewReader(file)

	var count uint64
	var line string
	for err == nil {
		line, err = r.ReadString('\n')
		if len(line) > 0 {
			count = ctr.Process(strings.Trim(line, "\n"))
		}
	}
	fmt.Printf("number of unique ipV4 addresses in the file %s is %d\n", filename, count)
}
