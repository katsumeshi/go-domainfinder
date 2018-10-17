package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transfroms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets " + otherWord,
	otherWord + "hq",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transfroms[rand.Intn(len(transfroms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
