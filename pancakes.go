package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"playground/kata/pancake2/waiters"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	numStacks, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	waiter := waiters.New()
	for i := 0; i < numStacks; {
		scanner.Scan()
		i++
		flips := waiter.ArrangeStack(scanner.Text())
		fmt.Println("Case #" + strconv.Itoa(i) +": " + strconv.Itoa(flips))
	}
}
