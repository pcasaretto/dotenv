package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/pcasaretto/dotenv/internal/parser"
)

func main() {
	var filename = flag.String("file", ".env", "file to load variables from")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s command [arg ...]:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Error opening file %s\n\t%s\n", *filename, err)
	}

	env, err := parser.Parse(f)
	if err != nil {
		log.Fatalf("Error parsing file %s\n\t%s\n", *filename, err)
	}
	env = append(env, os.Environ()...)

	cmdArgs := flag.Args()
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = env

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
