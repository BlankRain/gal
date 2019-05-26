package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/BlankRain/gal/exec"
	"github.com/BlankRain/gal/repl"
)

var f = flag.String("f", "", "fileName")
var r = flag.Bool("r", false, "start repl")

func main() {
	flag.Parse()
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if *f != "" {
		exec.Run(*f)
	}
	if *r {
		fmt.Printf("Hello %s! This is GAL REPL!\n", user.Username)
		fmt.Printf("Feel free  to type in commands \n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
