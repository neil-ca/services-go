package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "stranger", "your wonderful name")
var age = flag.Int("age", 0, "your age")
// go provides a basic library -> flag. It refers to the command-line flags
// ./flagExample -h -> "stranger"
// ./flagExample -name neil -age 21 -> Hello neil, .....
func main() {
	flag.Parse()
	log.Printf("Hello %s (%d years), Welcome to the command line", *name, *age)
}