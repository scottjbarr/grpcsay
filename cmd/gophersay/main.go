package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/scottjbarr/grpcsay"
)

func main() {
	msg := ""

	if len(os.Args) > 1 {
		msg = strings.Join(os.Args[1:], " ")
	}

	if len(msg) == 0 {
		// no data so trying stdin
		r := bufio.NewReader(os.Stdin)

		b, err := ioutil.ReadAll(r)
		if err != nil {
			panic(err)
		}

		msg = string(b)
	}

	msg = strings.TrimRight(msg, " \n")

	out := grpcsay.Say(msg)

	fmt.Printf("%v\n", out)
}
