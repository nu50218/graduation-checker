package main

import (
	"flag"
	"log"

	"github.com/nu50218/graduation-checker/checkers"
)

var target string

func init() {
	flag.StringVar(&target, "target", "", "e.g. suuri")
	flag.Parse()
}

func main() {
	cs, exists := checkers.GetCheckers(target)
	if !exists {
		log.Fatalf("target not exist: \"%s\"", target)
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("file not specified")
	}

	if err := cs.RunFile(args[0]); err != nil {
		log.Fatal(err)
	}
}
