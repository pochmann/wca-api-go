package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInt32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}

func waitForEnter(message string) {
	fmt.Print(message, " (waiting, hit Enter)")
	bufio.NewReader(os.Stdin).ReadLine()
}
