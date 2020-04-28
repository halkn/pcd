package main

import (
	"log"
	"os"

	"github.com/halkn/pcd/pcd"
)

func main() {

	err := pcd.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
