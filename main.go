package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	version = "1.0.0"
)

// CLI is Struct for stdio.
type CLI struct {
	outStream, errStream io.Writer
}

// Run a cli tools.
func (cli *CLI) Run(args []string) int {

	var (
		showVersion bool
	)

	flg := flag.NewFlagSet(args[0], flag.ExitOnError)
	flg.SetOutput(cli.errStream)
	flg.BoolVar(&showVersion, "v", false, "show version")
	if err := flg.Parse(args[1:]); err != nil {
		fmt.Fprintf(cli.outStream, "%s\n", err)
		os.Exit(1)
	}

	if showVersion {
		fmt.Fprintf(cli.outStream, "%s\n", "pcd's version is "+version)
		os.Exit(0)
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		os.Exit(1)
	}

	dirNames := strings.Split(filepath.ToSlash(pwd), "/")
	for idx := range dirNames {
		path := filepath.FromSlash(strings.Join(dirNames[0:idx+1], "/"))
		if path == "" {
			fmt.Fprintln(cli.outStream, "/")
		} else if strings.HasSuffix(path, ":") {
			fmt.Fprintln(cli.outStream, strings.ReplaceAll(path, ":", ":\\"))
		} else {
			fmt.Fprintln(cli.outStream, path)
		}
	}
	return 0
}

func main() {

	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}

	os.Exit(cli.Run(os.Args))
}
