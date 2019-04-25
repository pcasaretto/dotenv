package main

import (
	"flag"
	"fmt"
	"io"
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

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("Failed opening stdin")
	}
	go io.Copy(stdinPipe, os.Stdin)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed opening stdout")
	}
	go io.Copy(os.Stdout, stdoutPipe)
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Failed opening stderr")
	}
	go io.Copy(os.Stderr, stderrPipe)
	env = append(env, os.Environ()...)
	cmd.Env = env
	cmd.Run()
}
