// Package pcd ...
package pcd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	version = "0.0.1"
)

// Run the pcd.
func Run(argv []string, outStream, errStream io.Writer) error {
	var showVersion bool

	flg := flag.NewFlagSet("pcd", flag.ExitOnError)
	flg.SetOutput(errStream)

	flg.BoolVar(&showVersion, "version", false, "show version")
	if err := flg.Parse(argv[:]); err != nil {
		return fmt.Errorf("args parse error: %s", err)
	}

	if showVersion {
		return printVersion(outStream)
	}

	return printPathList(outStream)
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintln(out, "pcd's version is", version)
	if err != nil {
		return fmt.Errorf("print err version: %s", err)
	}
	return nil
}

func printPathList(out io.Writer) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error occurred While getting Current Directory : %s", err)
	}
	dirNames := strings.Split(filepath.ToSlash(wd), "/")
	lists := ""
	for idx := range dirNames {
		path := filepath.FromSlash(strings.Join(dirNames[0:idx+1], "/"))
		if path == "" {
			lists = "/\n"
		} else if strings.HasSuffix(path, ":") {
			lists = strings.ReplaceAll(path, ":", ":\\") + "\n"
		} else {
			lists += path + "\n"
		}
	}

	_, err = fmt.Fprint(out, lists)
	if err != nil {
		return fmt.Errorf("print err version: %s", err)
	}
	return nil
}
