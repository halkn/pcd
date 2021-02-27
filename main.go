package main

import (
	"log"
	"os"

	"github.com/halkn/pcd/pcd"
)

func main() {

	if err := pcd.Run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
